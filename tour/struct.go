package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func main() {

	y := Student{"yudamao", 21}

	fmt.Println(y)

}