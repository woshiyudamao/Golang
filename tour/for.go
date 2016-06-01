package main

import (
	"fmt"
)

func for1() {
	//基本的 for 循环
	s := 0
	for i := 1; i <= 1000; i++ {

		s += i
	}
	fmt.Printf("基本的for 循环 ret: %d \n", s)
}

func for2() {
	//去掉初始化语句和后置语句

	sum := 0
	for ; sum <= 1000; {
		sum += 1
	}
	fmt.Printf("去掉初始化语句和后置语句 循环 ret: %d \n", sum)
}

func for3() {
	//在2的基础上省区分号 类似于while 循环

	sum := 0
	for sum <= 1000 {
		sum += 1
	}

	fmt.Printf("在2的基础上省区分号 类似于while 循环 ret: %d \n", sum)
}

func main() {
	for1()
	for2()
	for3()

}
