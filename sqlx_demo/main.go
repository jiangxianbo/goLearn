package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

type user struct {
	Id   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_learn?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return nil
}

func queryRow() {
	sqlStr := "select `id`, `name`, `age` from user where `id` = ?"
	var u user
	err := db.Get(&u, sqlStr, 5)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(u)
}

func queryMultiRows() {
	var uList []user
	sqlStr := "select id, name, age from user"
	err := db.Select(&uList, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", uList)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB failed, err:%v\n", err)
		return
	}
	queryRow()
	queryMultiRows()
}
