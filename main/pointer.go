package main

import "fmt"

func pointer() {
	//var i1 = 5
	//fmt.Printf("这个数字%d,内存地址:%p\n", i1, &i1)
	//var intP *int
	//intP = &i1
	//fmt.Printf("这个内存地址:%p的值为:%d\n", intP, *intP)
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}
