package main

import "fmt"

// 递归：自己调用自己
// 适用适合处理问题相同、问题规模越来越小的场景
// 明确的退出条件
// 阶乘

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	return n * f(n-1)
}

// 上台阶问题
// n个台阶，一次可以走一步，也可以走两步，有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	//ret := f(5)
	//fmt.Println(ret)

	ret := taijie(4)
	fmt.Println(ret)
}
