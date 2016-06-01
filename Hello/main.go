package main

import (
	"fmt"
	"crypto/md5"
	"encoding/json"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	data := []byte("123")
	fmt.Printf("%x \n", md5.Sum(data))

	m := Message{
		Name: "Alice",
		Body: "Hello",
		Time: 1294706395881547000,
	}

	p := m

	fmt.Printf("%T", p)
	fmt.Println(json.Marshal(m))
}
