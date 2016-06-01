package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	//牛顿法求平方根
	z := 10.0
	ct := 0 //记录循环了多少次
	for math.Abs(z * z - x) > 1e-15 {
		z = z - (z * z - x) / ( 2 * z)
		ct += 1
	}
	return z, ct
}

func main() {
	x := 2.0
	a, b := Sqrt(x)
	fmt.Printf("牛顿法 Sqrt（%v）,循环  %v 次： %v,math.Sqrt(%v) : %v", x, b, a, x, math.Sqrt(x))

}
