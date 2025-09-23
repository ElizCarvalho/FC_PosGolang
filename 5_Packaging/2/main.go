package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("oi")
	fmt.Println(uuid.New().String())
}
