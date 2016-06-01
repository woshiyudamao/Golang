package main

import ("fmt"
	"math"
)
/*
Go 没有类。然而，仍然可以在结构体类型上定义方法。
方法接收者 出现在 func 关键字和方法名之间的参数中。

可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
但是，不能对来自其他包的类型或基础类型定义方法。

有两个原因需要使用指针接收者。首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。其次，方法可以修改接收者指向的值。
*/
type Student struct {
	age,sex int
	name string
	address string
	email string
}

type MyFloat float64

func (f MyFloat) Abs() float64 { //这里不是指针是因为不需要对值进行改变
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (t * Student) Show() () {

	fmt.Printf("我是 %v,今年 %v 岁 .家住 %v \n",t.name,t.age,t.address )
	t.age++   //这里的改变会影响
	return
}

func main() {


	s:=Student{18,1,"YuDaMao","湖北武汉","example@qq.com"}
	s.Show()
	fmt.Printf("明年就%v岁啦～\n",s.age)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
