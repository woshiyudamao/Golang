package main

import "fmt"
/*
Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 nil 。

var p *int
& 符号会生成一个指向其作用对象的指针。

i := 42
p = &i
* 符号表示指针指向的底层的值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“非直接引用”。

与 C 不同，Go 没有指针运算。
*/
func main() {

	a, b := 10, 20
	fmt.Printf("&a=%v,&b=%v \n", &a, &b)
	*&a = 20

	*&b = 10

	fmt.Printf("a=%v,b=%v", a, b)

}
