package main

import (
	"encoding/json"
	"fmt"
)

// 只有struct中支持导出的field才能被JSON package序列化,即首字母大写的field
// Struct tag可以决定Marshal和Unmarshal函数如何序列化和反序列化数据
// 1.JSON object中的name一般都是小写,我们可以通过struct tag来实现
// 2.使用omitempty可以告诉Marshal函数如果field的值是对应类型的zero-value，那么序列化之后的JSON object中不包含此field
// 3.Struct tag “-” 表示跳过指定的filed
type Message struct {
	Name string `json:"name"`
	Body string
	Time int64 `json:"-"`
	Id   int   `json:"id,omitempty"`
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000, 0}
	b, _ := json.Marshal(m)
	fmt.Println(string(b))

	var s Message
	json.Unmarshal(b, &s)
	fmt.Println(s)

	fmt.Println("---------------------------")

	j := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(j, &f)
	ff := f.(map[string]interface{})
	fmt.Println(ff["Name"])
	fmt.Println(ff["Age"])
	fmt.Println(ff["Parents"].([]interface{})[0])

	for k, v := range ff {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array", vv)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	fmt.Println("---------------------------")

	var fm FamilyMember
	json.Unmarshal(j, &fm)
	fmt.Println(fm)
}
