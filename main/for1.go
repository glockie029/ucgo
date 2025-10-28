package main

import (
	"fmt"
	"strings"
)

func for1() {
	var p int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d+%d=", p, j)
			p += j
			fmt.Println(p)
		}
	}
}
func bunicode() {
	str := "Go is Beautiful language!"
	fmt.Printf("The length of str is %d\n", len(str))
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on postion %d is :%c\n", ix, str[ix])
	}
}
func simplefor() {
	for i := 0; i < 26; i++ {
		fmt.Println(strings.Repeat("G", i))
	}
}
func bm() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d的补码%b,10进制为%d\n", i, ^i, int(^i))
	}
}
func FIZZ() {
	for i := 0; i < 100; i++ {
		switch true {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FIZZBUZZ: ", i)
		case i%3 == 0:
			fmt.Println("FIZZ: ", i)
		case i%5 == 0:
			fmt.Println("BUZZ: ", i)
		default:
		}
	}
}
func suqare() {
	var heigh string
	//var width string
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			heigh = strings.Repeat("*", j)
		}
		fmt.Println(heigh)
	}

}
func foreach() {
	//str := "I'm your father"
	//for pos, char := range str {
	//	fmt.Printf("Character on postion %d is : %c\n", pos, char)
	//}
	//str1 := "我爱中国"
	//for pos, char := range str1 {
	//	fmt.Printf("%s的第%d个字符是: %c\n", str1, pos, char)
	//}
	//for index, rune := range str1 {
	//	fmt.Printf("%-2d %d %U %c % X\n", index, rune, rune, rune, rune)
	//}
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
