package main

import (
	"bytes"
	"fmt"
	"time"
)

func ListPractise() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1的第%d个是%d\n", i, arr1[i])
	}
	for k, v := range arr1 {
		fmt.Printf("arr1的第%d个是%d\n", k, v)
	}
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	var arr2 = arr1
	//fmt.Println(arr2)
	arr2[0] = 99
	fmt.Println(arr2)
	fmt.Println(arr1)
	const n = 50
	var fibs [n]uint64
	if n > 0 {
		fibs[0] = 0
	}
	if n > 1 {
		fibs[1] = 1
	}
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	fmt.Println(fibs)
}

const (
	HIGHT  = 20
	WEIGHT = 10
)

type pixel int

var screen [HIGHT][WEIGHT]pixel

func mul_array() {

	for y := 0; y < HIGHT; y++ {
		fmt.Println()
		for x := 0; x < WEIGHT; x++ {
			fmt.Printf("%d", screen[y][x])
		}
	}

}

func slicefib(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	fmt.Println(a)
}

func bytes_test() {
	start := time.Now()
	var s string
	for i := 0; i < 10000; i++ {
		s += "a"
	}
	fmt.Printf("字符串拼接耗时: %s\n", time.Since(start))
	start = time.Now()
	var b bytes.Buffer
	for i := 0; i < 10000; i++ {
		b.WriteString("a")
	}
	_ = b.String()
	fmt.Printf("字符串拼接耗时:%s\n", time.Since(start))
}
