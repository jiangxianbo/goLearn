package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "Hello沙河"
	count := 0
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	s2 := "how do you do"
	s3 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
	}
	fmt.Println(m1)

	// 回文判断
	ss := "山西运煤车煤运西山"
	r := make([]rune, 0, len(ss))
	fmt.Println(r)
	for _, v := range ss {
		r = append(r, v)
	}
	fmt.Println(r)
	for i := 0; i < len(r)-1-i; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
