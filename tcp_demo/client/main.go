package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client

func main() {
	// 1. 与server简历连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("Dial 127.0.0.1:20000 failed, err:%v\n", err)
	}

	// 2. 发送数据
	reader := bufio.NewReader(os.Stdin)
	var tmp [128]byte
	for {
		fmt.Print("说：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))

		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}

	conn.Close()
}
