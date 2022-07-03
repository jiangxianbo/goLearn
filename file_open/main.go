package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFromFile() {
	// 打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file faild, err:%v", err)
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读取文件
	//var tmp = make([]byte, 128)
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file fail, err:%v", err)
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

// bufio读取文件
func readFromFileByBufio() {
	// 打开文件
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file faild, err:%v", err)
	}
	defer fileObj.Close()
	// 创建一个从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	readFromFile()
	//readFromFileByBufio()
}
