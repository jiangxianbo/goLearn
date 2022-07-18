package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_learn"
	db, err := sql.Open("mysql", dsn) // 不会校验
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
