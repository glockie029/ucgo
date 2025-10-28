package main

import "fmt"

const (
	Red = iota
	Green
	Yellow
	Gray
	Silver
	Brown
	Blue = 6
	White
	Black
)

func sswitch() {
	var months int = 5
	switch months {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Printf("夏季")
	case 9, 10, 11:
		fmt.Print("秋季")
	case 12, 1, 2:
		fmt.Printf("冬季")
	default:
		fmt.Printf("非法月份!")
	}
}

func ioatf() {

}
