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
	ctxx, cannel := context.WithTimeout(ctx, 5*time.Second)
	defer cannel()

	_, err := redisDB.Ping(ctxx).Result()
	if err != nil {
		fmt.Printf("redis connect failed, err%s\n", err)
		return nil, err
	}
	fmt.Println("redis connect success！")
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
	stringIntExample()
}
