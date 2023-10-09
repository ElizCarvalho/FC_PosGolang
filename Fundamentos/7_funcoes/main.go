package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(3, 5))
	result, err := sum(30, 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func sum(num1, num2 int) (int, error) {
	if num1+num2 >= 50 {
		return 0, errors.New("a soma é maior que 50")
	}
	return num1 + num2, nil
}
