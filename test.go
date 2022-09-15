package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo Student
	think := "speak"
	fmt.Println(peo.Speak(think))
}

// 29 28 28

// 10 1 2 3
// 20 0 2 2
