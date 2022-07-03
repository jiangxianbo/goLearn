package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开要操作的文件
	fileObj, err := os.OpenFile("./a.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}

	// 创建临时文件
	tmpFile, err := os.OpenFile("./a.tmp", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer tmpFile.Close()

	// 读源文件，写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read file failed, err:%v\n", err)
		return
	}
	// 写入临时文件
	tmpFile.Write(ret[:n])

	// 在写入要插入的内容
	var s []byte
	s = []byte{'c'}
	tmpFile.Write(s)

	// 把源文件后续内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tmpFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v\n", err)
			return
		}
		tmpFile.Write(x[:n])
	}
	// 源文件后续的也写入临时文件
	fileObj.Close()
	tmpFile.Close()
	os.Rename("./a.tmp", "./a.txt")
}
