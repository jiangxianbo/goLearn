package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数校验
	// 0.1 传进来的data参数必须是指针类型（因为需要赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针（因为配置文件中有各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}
	// 1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b) // 将字节类型转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	//fmt.Printf("%#v\n", lineSlice)
	// 2. 一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉首位空格
		line = strings.TrimSpace(line)
		// 空行处理
		if len(line) == 0 {
			continue
		}
		// 2.1 注释忽略
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[]就是节(section)
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 把这一行首尾的[]去掉，取到中间内容把首尾空格去掉
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 找到了对应的嵌套结构体， 把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s, 对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3 如果不是[]就是需要数据
			// 2.3.1 等号分割
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntx error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2.3.2 根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)

			sValue := v.Elem().FieldByName(structName) // 结构体值信息
			sType := sValue.Type()                     // 嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 2.3.3 遍历嵌套结构体每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i) // tag信息存储在类型信息中
				fileType = field
				if field.Tag.Get("ini") == key {
					// 找到对应字段
					fieldName = field.Name
					break
				}
			}
			// 2.3.4 如果key = tag，给这个字段赋值
			// 2.3.4.1 根据fieldName 去去取出这个字段
			if len(fieldName) == 0 {
				// 在结构体重找不到对应字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			// 2.3.4.2 对其赋值
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			}
		}
	}
	return

}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
