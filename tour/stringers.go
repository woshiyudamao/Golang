package main

import (

	"fmt"
)
type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("我是 %v ,今年 %v 岁！\n",p.name,p.age)
}



type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v",ip[0],ip[1],ip[2],ip[3])

}




func main() {
	p:=Person{"Yu",18}
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(p,a, z)
	/*
	练习：Stringers
	让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。
	例如，IPAddr{1, 2, 3, 4} 应当输出 "1.2.3.4"。
	 */
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}

}