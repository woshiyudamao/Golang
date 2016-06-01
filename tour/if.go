package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//if 的便捷语句

	/*
	跟 for 一样， if 语句可以在条件之前执行一个简单语句。
 	由这个语句定义的变量的作用域仅在 if 范围之内。
	 */
	if v := math.Pow(x, n); v < lim {
		//lim 是限制的最大值
		return v
	}
	return lim
}

func main() {

	fmt.Println("begin")
	n := 10
	if n < 100 {
		//if common use
		fmt.Println("n less than 100")
		fmt.Println(math.Sqrt(float64(n)))

	}
	fmt.Println(pow(2, 10, 1000))

	if 10 < 2 {

	} else {
		fmt.Println("mmm_mmm")
	}

}
