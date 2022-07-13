package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 对于GET请求，参数放在URL(query param), 请求体无数数据
	fmt.Println(r.URL)
	queryPram := r.URL.Query() // 自动识别URL中query param
	name := queryPram.Get("name")
	age := queryPram.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.URL.Query())
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
