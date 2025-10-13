package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	var p fastjson.Parser
	jsonData := `{"foo": "bar", "num": 42, "bool": true, "arr": [1, 2, 3], "obj": {"key": "value"}}`

	v, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("1- foo: %s\n", v.GetStringBytes("foo"))
	fmt.Printf("2-num: %d\n", v.GetInt("num"))
	fmt.Printf("3-bool: %t\n", v.GetBool("bool"))
	fmt.Printf("4-arr: %v\n", v.GetArray("arr"))

	fmt.Println("--------------------------------")
	a := v.GetArray("arr")
	for i, value := range a {
		fmt.Printf("Index: %d, Value: %s\n", i, value)
	}
	fmt.Println("--------------------------------")
	fmt.Printf("5-obj: %v\n", v.GetObject("obj"))
	fmt.Printf("6-obj.key: %s\n", v.GetStringBytes("obj.key"))
	fmt.Println(string(v.MarshalTo([]byte{})))
}
