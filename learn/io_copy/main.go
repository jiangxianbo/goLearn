package main

import (
	"fmt"
	"io"
	"os"
)

// 借助io.copy() 实现一个拷贝文件函数

func CopyFile(dstname, srcName string) (written int64, err error) {
	// 以读的方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", err)
		return
	}
	defer src.Close()
	// 以写|创建的方式打开没有标文件
	dst, err := os.OpenFile("dstName", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed, err:%v.\n", err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)

}

func main() {
	_, err := CopyFile("dst.txt", "src.txt")
	if err != nil {
		fmt.Printf("copy file failed, err:%v.\n", err)
		return
	}
	fmt.Println("copy done!")
}
