package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP client

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Printf("连接服务器失败, err:%v\n", err)
		return
	}
	defer socket.Close()
	// 发送数据
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("输入:")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))

		// 收到回复的数据
		n, addr, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Printf("redv reply msg failed, err:%v\n", err)
			return
		}
		fmt.Printf("回复地址：%v, msg:%v\n", addr, string(reply[:n]))
	}

}
