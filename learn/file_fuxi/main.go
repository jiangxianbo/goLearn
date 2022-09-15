package main

import (
	"fmt"
	"os"
)

func f11() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./mylogger.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close()
}
func main() {

}
