package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_learn"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func transactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("db begion failed, err:%v\n", err)
		return
	}

	sqlStr1 := "update user set age = age - 2 where id = ?"
	sqlStr2 := "update user set age = age + 2 where id = ?"
	_, err = tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Printf("commit failed, err:%v\n", err)
		return
	}
	fmt.Println("执行成功！")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err: %v\n", err)
	}
	fmt.Println("连接数据库成功！")

	transactionDemo()
}
