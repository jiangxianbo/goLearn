package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./a.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	fileObj.Seek(3, 0)
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	fmt.Println(string(ret[:n]))
}
