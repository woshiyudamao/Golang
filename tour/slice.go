package main


/*
对 slice 切片
slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

表达式

s[lo:hi]
表示从 lo 到 hi-1 的 slice 元素，含前端，不包含后端。因此

s[lo:lo]
是空的，而

s[lo:lo+1]
有一个元素。
 */


import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])

	a := new([10]int)
	a[1] = 10
	b := make([]int, 0, 5)
	b = append(b, 10) //向 slice 的末尾添加元素是一种常见的操作，因此 Go 提供了一个内建函数 append 。 内建函数的文档对 append 有详细介绍。
	b = append(b, 10)
	b = append(b, 10)
	b = append(b, 10)
	b = append(b, 210)
	b = append(b, 2110)
	fmt.Println(a)
	fmt.Println(cap(b))
	//b[100]=1

	/*
	nil slice
	slice 的零值是 nil 。
	一个 nil 的 slice 的长度和容量是 0。
	 */


	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}





	//range 遍历数组

	for t1, t2 := range b {
		fmt.Printf("%v = %d\n", t1, t2)
	}

	for _, t2 := range b {
		fmt.Printf("%d\n", t2)
	}

}
