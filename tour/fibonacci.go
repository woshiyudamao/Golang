package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
//斐波那契数列：0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144
func fibonacci() func() int {
	//自己写的
	f1, f2, t := 0, 0, 0
	return func() int {
		if t == 1 {
			f1 = 1
		} else {
			f1, f2 = f1 + f2, f1 //这种赋值是立即产生的
		}
		t++
		return f1
	}
}

func fibonacci2(num int) int {
	//递归实现  没有显示第一个0
	if num < 2 {
		return 1
	}
	return fibonacci2(num - 1) + fibonacci2(num - 2)
}

func fib() func() int {
	//Go的闭包实现，摘自golang官网 没有显示第一个0
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func main() {
	f := fibonacci()
	f2 := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println("------------------------------------")
	for i := 0; i < 10; i++ {
		fmt.Println(f2())
	}

}
