package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	os := runtime.GOOS

	fmt.Printf("当前系统为： %v \n", os)

	switch os {

	case "linux" :

		fmt.Println("Linux!")
	//runtime.Breakpoint()
	case "windows":

		fmt.Println("windows!")



	}



	/*
	没有条件的 switch 同 switch true 一样。
	这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。
	*/

	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("早上好")
	case t.Hour() < 17 :
		fmt.Println("下午好")
	default :
		fmt.Println("晚上好")

	}

}