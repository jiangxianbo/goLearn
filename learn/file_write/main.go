package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func write() {
	fileObj, err := os.OpenFile("./xx_write.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file fail,err:%v", err)

	}
	defer fileObj.Close()
	// 写入字节切片数据
	fileObj.Write([]byte("write byte\n"))
	// 写入字符串
	fileObj.WriteString("写入字符串\n")
}

func writeByBufio() {
	fileObj, err := os.OpenFile("./xx_bufio.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello world\n") // 将数据先写入缓存
	wr.Flush()                      // 将缓存中的内容写入文件
}

func writeByIoutil() {
	str := "hello 沙河"
	err := ioutil.WriteFile("./xx_ioutil.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

// 打开文件写内容
func main() {
	//write()
	//writeByBufio()
	writeByIoutil()
}
