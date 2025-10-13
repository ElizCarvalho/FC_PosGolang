package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	var p fastjson.Parser
	jsonData := `{"user": {"name": "John Doe", "age": 30, "city": "New York"}}`
	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	userJSON := value.GetObject("user").String()
	var user User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User name: %s\n", user.Name)
	fmt.Printf("User age: %d\n", user.Age)
	fmt.Printf("User city: %s\n", user.City)
}
