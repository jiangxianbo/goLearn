package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	left := 50
	for _, user := range users {
		//for _, c := range user {
		//	switch c {
		//	case 'e', 'E':
		//	case 'i', 'I':
		//
		//	}
		//}
		gold := 0
		enum := strings.Count(user, "e")
		Enum := strings.Count(user, "E")
		inum := strings.Count(user, "i")
		Inum := strings.Count(user, "I")
		onum := strings.Count(user, "o")
		Onum := strings.Count(user, "O")
		unum := strings.Count(user, "u")
		Unum := strings.Count(user, "U")
		if enum > 0 || Enum > 0 {
			gold += (enum + Enum) * 1
		}
		if inum > 0 || Inum > 0 {
			gold += (inum + Inum) * 2
		}
		if onum > 0 || Onum > 0 {
			gold += (onum + Onum) * 3
		}
		if unum > 0 || Unum > 0 {
			gold += (unum + Unum) * 4
		}
		distribution[user] = gold
		left -= gold

	}
	fmt.Println(distribution)
	return left
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
