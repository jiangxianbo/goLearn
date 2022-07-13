package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client

func main() {
	//resp, err := http.Get("http://127.0.0.1:9090/xxx/?name=sb&age=18")
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}
	data := url.Values{} // url values
	urlObj, _ := url.Parse("http://127.0.0.1:9090/xxx/")
	data.Set("name", "周琳")
	data.Set("age", "9000")
	queryStr := data.Encode() // URL encode之后的URl
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close() // 一定要记得关闭 resp.Body
	// 发送请求
	// 从resp中把服务器返回的数据读出来
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
