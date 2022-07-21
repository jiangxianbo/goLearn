package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置数据库连接池的最大空闲连接数
	db.SetMaxIdleConns(10)
	return nil
}

func queryOne(id int) (ret user) {
	var u user
	// 1.单条记录
	sqlStr := "select id, name, age from user where id = ?;"
	// 2.执行并拿到记录
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		panic(err)
	}
	// 打印结果
	return u
}

func queryMore(n int) {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		panic(err)
	}
	// 一定要关闭连接
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			panic(err)
		}
		fmt.Println(u)
	}
}

// 插入数据
func insert() {
	sqlStr := "insert into user(`name`, `age`) values (?,?)"
	ret, err := db.Exec(sqlStr, "测试插入", 18)
	if err != nil {
		panic(err)
	}
	id, err := ret.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("insert success, the id is %d\n", id)
}

func updateRow(id int, age int) {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		panic(err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("更新了%d行数据\n", n)
}

func delUser(id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		panic(err)
	}
	n, err := ret.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("del success, affect row: %d\n", n)
}

// 预处理
func prepareInsert() {
	sqlStr := "insert into user(`name`, `age`) values(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var m = map[string]int{
		"test001": 1,
		"test002": 2,
		"test003": 3,
		"test004": 4,
	}
	for s, i := range m {
		_, err = stmt.Exec(s, i)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}

	//a := queryOne(10)
	//queryMore(1)
	//insert()
	//updateRow(1, 110)
	//delUser(1)
	prepareInsert()
}
