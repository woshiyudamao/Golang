package main

import (
	"fmt"
)

var global int = 0

func main() {
	/*
	defer 语句会延迟函数的执行直到上层函数返回。
	 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
	延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
	*/
	defer fmt.Println("我是defer语句")

	fmt.Println("Hello World")

	defer fmt.Printf("全局变量 Global = %v \n", global) //验证了参数会立刻生成

	global = 10

	for i := 0; i < 10; i++ {
		//   last in first out
		defer fmt.Println(i)
	}

}