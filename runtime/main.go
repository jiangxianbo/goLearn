package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.Caller failed, err:%v\n")
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	f1()
}
