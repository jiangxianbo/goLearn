package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	redisDB *redis.Client
	ctx     = context.Background()
)

type RedisSingleObj struct {
	host     string
	pwd      string
	DB       int
	PoolSize int
}

func (r RedisSingleObj) initRedis() (*redis.Client, error) {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     r.host,
		Password: r.pwd,
		DB:       r.DB,
		PoolSize: r.PoolSize, // 连接池大小
	})
	fmt.Printf("Connecting Redis : %v\n", r.host)

	ctxx, cannel := context.WithTimeout(ctx, 5*time.Second)
	defer cannel()

	res, err := redisDB.Ping(ctxx).Result()
	if err != nil {
		fmt.Printf("redis connect failed, err%s\n", err)
		return nil, err
	}
	fmt.Printf("Connect Successful! Ping => %v\n", res)
	return redisDB, nil
}

// setGetExample redis string
func setGetExample() {
	// 1.Set 设置 key 如果设置-1永不过期
	err := redisDB.Set(ctx, "score", 100, time.Second*60).Err()
	if err != nil {
		fmt.Printf("Set score failed, err:%v\n", err)
		panic(err)
	}

	// 2.Get 获取值
	val1, err := redisDB.Get(ctx, "score").Result()
	if err != nil {
		fmt.Printf("Get score failed, err:%v\n", err)
		panic(err)
	}
	fmt.Printf("ret-> score:%v\n", val1)

	// Get 获取一个不存在的值
	val2, err := redisDB.Get(ctx, "test").Result()
	fmt.Printf("Get 获取一个不存在的值, val2:%v\n", val2)
	if err == redis.Nil {
		fmt.Println("[ERROR] - Key [test] not exist")
		fmt.Printf("Get 获取一个不存在的值, err=>%v\n", err)
	} else if err != nil {
		fmt.Printf("Get name failed, err:%v\n", err)
		panic(err)
	}

	// Exists() 方法检测某个key是否存在
	n, _ := redisDB.Exists(ctx, "name").Result()
	if n > 0 {
		fmt.Println("name key 存在！")
	} else {
		fmt.Println("name key 不存在！")
		redisDB.Set(ctx, "name", "jxb", time.Second*60)
	}

	// 3. SetNX 当不存在key时将进行设置该可以并设置其过期时间
	val3, err := redisDB.SetNX(ctx, "username", "jiangxianbo", 0).Result()
	if err != nil {
		fmt.Printf("set username failed, err:%v\n", err)
		panic(err)
	}
	fmt.Printf("val3 -> username: %v\n", val3)

	// 4. keys() 根据正则获取keys, DBSize() 查看当前数据库key的数量
	keys, _ := redisDB.Keys(ctx, "*").Result()
	num, err := redisDB.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("All Keys : %v, Keys number : %v \n", keys, num)

	// 5.Type() 获取一个key对应值得类型
	vType, err := redisDB.Type(ctx, "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("username key type : %v\n", vType)

	// 6.Expire() 设置某时间段（time.Duration）后过期，ExpireAt()某时间点（time.time）过期失效
	val4, _ := redisDB.Expire(ctx, "name", time.Minute*2).Result()
	if val4 {
		fmt.Println("name 过期时间设置成功", val4)
	} else {
		fmt.Println("name 过期时间设置失败", val4)
	}
	val5, _ := redisDB.ExpireAt(ctx, "username", time.Now().Add(time.Minute*2)).Result()
	if val5 {
		fmt.Println("username 过期时间设置成功", val5)
	} else {
		fmt.Println("username 过期时间设置失败", val5)
	}

	// 7.TTL()与PTTL()方法可以获取某个键的剩余有效期
	userTTL, _ := redisDB.TTL(ctx, "user").Result() // 获取key的有效时间
	usernameTTL, _ := redisDB.PTTL(ctx, "username").Result()
	fmt.Printf("user TTL : %v, username TTL : %v\n", userTTL, usernameTTL)

	// 8.Del()：删除缓存、FlushDB():清空当前数据
	// 当通配符匹配的key的数量不多时，可以使用Keys()得到所有的key在使用Del命令删除。
	num, err = redisDB.Del(ctx, "user", "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Del() : ", num)
	// 如果key的数量非常多的时候，我们可以搭配使用Scan命令和Del命令完成删除。
	iter := redisDB.Scan(ctx, 0, "user*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println(iter.Val())
		err := redisDB.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	// 9.清空当前数据库，因为连接的是索引为0的数据库，所以清空的就是0号数据库
	flag, err := redisDB.FlushDB(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("FlushDB() : ", flag)
}

// stringIntExample redis string/int
func stringIntExample() {
	// 设置字符串类型key
	err := redisDB.Set(ctx, "hello", "hello world!", 0).Err()
	if err != nil {
		panic(err)
	}
	// GetRange：字符串截取
	// 注：即使key不存在，调用GetRange()也不会报错，只是返回的截取结果是空"",可以使用fmt.Printf("%q\n", val)来打印测试
	val1, _ := redisDB.GetRange(ctx, "hello", 1, 4).Result()
	fmt.Printf("key: hello, value: %v\n", val1) //截取到的内容为: ello

	// Append() 往字符串后面追加元素， 返回值是字符串的总长度
	length1, _ := redisDB.Append(ctx, "hello", "Go Programer").Result()
	val2, _ := redisDB.Get(ctx, "hello").Result()
	fmt.Printf("当前缓存key的长度为: %v，值: %v \n", length1, val2)

	// 设置整形的key
	err = redisDB.Set(ctx, "number", 1, 0).Err()
	if err != nil {
		panic(err)
	}
	// Incr()、IncrBy()都是操作数字，对数字进行增加的操作
	// Decr()、DecrBy()方法是对数字进行减的操作，和Incr正好相反
	// incr是执行原子加1操作
	val3, _ := redisDB.Incr(ctx, "number").Result()
	fmt.Printf("Incr -> key当前的值为: %v\n", val3) // 2
	// incrBy是增加指定的数
	val4, _ := redisDB.IncrBy(ctx, "number", 6).Result()
	fmt.Printf("IncrBy -> key当前的值为: %v\n", val4) // 8

	//	Strlen 也可以返回缓存key的长度
	length2, _ := redisDB.StrLen(ctx, "number").Result()
	fmt.Printf("number 值长度: %v\n", length2)
}

// listExample 列表list类型操作
func listExample() {
	// 插入指定值到list列表中，返回值是当前列表元素的数量
	// 使用 LPush() 方法将数据从左侧压入链表（先进先出），也可以从右侧压入链表对应的方法是RPush()
	count, _ := redisDB.LPush(ctx, "list", 1, 2, 3).Result()
	fmt.Println("插入到list集合中元素的数量: ", count)

	// LInsert() 在某个位置插入新元素
	// 在名为key的缓存项值为2的元素前面插入一个值，值为123 ， 注意只会执行一次
	_ = redisDB.LInsert(ctx, "list", "before", "2", 123).Err()
	// 在名为key的缓存项值为2的元素后面插入一个值，值为321
	_ = redisDB.LInsert(ctx, "list", "after", "2", 321).Err()

	// LSet() 设置某个元素的值
	// 下标是从0开始的
	val1, _ := redisDB.LSet(ctx, "list", 2, 256).Result()
	fmt.Println("是否成功将下标为2的元素值改成256: ", val1)

	// LLen() 获取链表元素个数
	length, _ := redisDB.LLen(ctx, "list").Result()
	fmt.Printf("当前链表的长度为: %v\n", length)

	// LIndex() 获取链表下标对应的元素
	val2, _ := redisDB.LIndex(ctx, "list", 2).Result()
	fmt.Printf("下标为2的值为: %v\n", val2)

	// 从链表左侧弹出数据
	val3, _ := redisDB.LPop(ctx, "list").Result()
	fmt.Printf("弹出下标为0的值为: %v\n", val3)

	// LRem() 根据值移除元素 lrem key count value
	n, _ := redisDB.LRem(ctx, "list", 2, 256).Result()
	fmt.Printf("移除了: %v 个\n", n)
}

// setExample 集合set操
func setExample() {
	// 集合元素缓存设置
	key := "Program"
	mem := []string{"C", "Golang", "C++", "C#", "Java", "Delphi", "Python", "Golang"}
	for _, s := range mem {
		redisDB.SAdd(ctx, key, s).Result()
	}

	// SCard() 获取集合元素个数
	total, _ := redisDB.SCard(ctx, key).Result()
	fmt.Println("golang集合成员个数: ", total)

	// SPop() 随机获取一个元素 （无序性，是随机的）
	val1, _ := redisDB.SPop(ctx, key).Result()
	// SPopN()  随机获取多个元素
	val2, _ := redisDB.SPopN(ctx, key, 2).Result()
	// SSMembers() 获取所有成员
	val3, _ := redisDB.SMembers(ctx, key).Result()
	fmt.Printf("随机获取一个元素: %v , 随机获取多个元素: %v \n所有成员: %v\n", val1, val2, val3)

	// SIsMember() 判断元素是否在集合中
	exists, _ := redisDB.SIsMember(ctx, key, "golang").Result()
	if exists {
		fmt.Println("golang 存在 Program 集合中.") // 注意:我们存入的是Golang而非golang
	} else {
		fmt.Println("golang 不存在 Program 集合中.")
	}

	// SUnion():并集, SDiff():差集, SInter():交集
	redisDB.SAdd(ctx, "setA", "a", "b", "c", "d")
	redisDB.SAdd(ctx, "setB", "a", "d", "e", "f")

	//并集
	union, _ := redisDB.SUnion(ctx, "setA", "setB").Result()
	fmt.Println("并集", union)

	//差集
	diff, _ := redisDB.SDiff(ctx, "setA", "setB").Result()
	fmt.Println("差集", diff)

	//交集
	inter, _ := redisDB.SInter(ctx, "setA", "setB").Result()
	fmt.Println("交集", inter)

	// 删除集合中指定元素(返回成功)
	n, _ := redisDB.SRem(ctx, "setB", "a", "f").Result()
	fmt.Println("已成功删除元素的个数: ", n)
}

// zsetExample zset有序集合
func zsetExample() {
	// 有序集合成员与分数设置
	// zSet类型需要使用特定的类型值*redis.Z，以便作为排序使用
	lang := []*redis.Z{
		&redis.Z{Score: 90, Member: "Golang"},
		&redis.Z{Score: 98, Member: "Java"},
		&redis.Z{Score: 95, Member: "Python"},
		&redis.Z{Score: 97, Member: "JavaScript"},
		&redis.Z{Score: 99, Member: "C/C++"},
	}
	// ZAdd() 插入ZSet类型
	num, err := redisDB.ZAdd(ctx, "language_rank", lang...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// ZIncrBy() 将ZSet中的某一个元素顺序值增加: 把Golang的分数加10
	newScore, err := redisDB.ZIncrBy(ctx, "language_rank", 10, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 根据分数排名取出元素:取分数最高的3个
	ret, err := redisDB.ZRangeWithScores(ctx, "language_rank", 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	fmt.Printf("zsetKey前3名热度的是: %v\n Top 3 的 Memeber 与 Score 是:\n", ret)
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	// ZRangeByScore() 通过分数返回有序集合指定区间内的成员
	// ZRevRangeByScore() 返回有序集中指定分数区间内的成员，分数从高到低排序
	// 此处表示取95~100分的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = redisDB.ZRangeByScoreWithScores(ctx, "language_rank", &op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	// 输出全部成员及其score分数
	fmt.Println("language_rank 键存储的全部元素:")
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// hashExample 哈希(hash)类型操作
// hash 是一个 string 类型的 field（字段） 和 value（值） 的映射表，hash 特别适合用于存储对象。
func hashExample() {
	// (1) HSet() 设置字段和值
	redisDB.HSet(ctx, "hmuser", "key1", "value1").Err()
	redisDB.HSet(ctx, "hmuser", []string{"key2", "value2"}).Err()
	redisDB.HSet(ctx, "hmuser", map[string]interface{}{"key4": "value4"})

	// (2) HMset():批量设置
	redisDB.HMSet(ctx, "hmuser", map[string]interface{}{"name": "WeiyiGeek", "age": 88, "address": "重庆"})

	// (3) HGet() 获取某个元素
	address, _ := redisDB.HGet(ctx, "hmuser", "address").Result()
	fmt.Println("hmuser.address -> ", address)

	// (4) HGetAll() 获取全部元素
	hmuser, _ := redisDB.HGetAll(ctx, "hmuser").Result()
	fmt.Println("hmuser :=> ", hmuser)

	// (5) HExists 判断元素是否存在
	flag, _ := redisDB.HExists(ctx, "hmuser", "address").Result()
	fmt.Println("address 是否存在 hmuser 中: ", flag)

	// (6) HLen() 获取长度
	length, _ := redisDB.HLen(ctx, "hmuser").Result()
	fmt.Println("hmuser hash 键长度: ", length)

	// (7) HDel() 支持一次删除多个元素
	count, _ := redisDB.HDel(ctx, "hmuser", "key4", "key3").Result()
	fmt.Println("删除元素的个数: ", count)
}

func main() {
	conn := &RedisSingleObj{
		host:     "127.0.0.1:6379",
		pwd:      "",
		DB:       0,
		PoolSize: 100,
	}
	_, err := conn.initRedis()
	if err != nil {
		fmt.Printf("[连接Redis失败], err：%s\n", err)
		return
	}
	defer redisDB.Close()
	//setGetExample()
	//stringIntExample()
	//listExample()
	//setExample()
	//zsetExample()
	hashExample()
}
