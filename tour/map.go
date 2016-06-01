package main

import (
	"fmt"
)
/*
map 映射键到值。

map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。
 */

type Vertex struct {
	Lat, Long float64
}

func main() {
	fmt.Println("map test")
	//var m map[string]int
	m := make(map[string]int)
	m["xx"] = 2
	fmt.Println(m == nil)

	m2 := map[string]Vertex{
		"Bell Labs" : {40.6, -74.3},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m2)



	//m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer") //删除元素
	fmt.Println("The value:", m["Answer"]) //读取到不存在的元素结果是元素的0值

	v, ok := m["Answer"] //双赋值检测是否存在 ok =true or false

	fmt.Println("The value:", v, "Present?", ok)

}