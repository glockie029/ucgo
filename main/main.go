package main

import (
	"PracticeGo/structpractice"
	"fmt"
	"reflect"
	"unicode/utf8"
)

func analyzeString(s string, name string) {
	fmt.Printf("--- 分析字符串: \"%s\" ---\n", s)
	byteCount := len(s)
	fmt.Printf("字节(byte)数量: %d\n", byteCount)
	runeCount := utf8.RuneCountInString(s)
	fmt.Printf("字符(rune)的数量: %d\n", runeCount)
	fmt.Printf("----------------------------\n")
}

//func lastSeven(s string) {
//	t := time.Now()
//}

func zZhen() {
	var i1 = 1
	fmt.Printf("%d的内存地址:%p", i1, &i1)
	var intA *int
	intA = &i1
	fmt.Println(intA, *intA)
	var i2 = &i1
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(*i2)
}
func main() {
	//printvalue()
	//closepack()
	//test68()
	//start := time.Now()
	//addhttp := MakeAddscheme("http://")
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Println(delta)
	//addhttps := MakeAddscheme("https://")
	//fmt.Println(addhttp("www.baidu.com"))
	//fmt.Println(addhttps("www.baidu.com"))
	//var result uint64 = 0
	//start := time.Now()
	//for i := 0; i < LIM; i++ {
	//	result = fibonacci(i)
	//	fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	//}
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	//ListPractise()
	//mul_array()
	//var list1 [9]int
	//var list2 []int = list1[2:5]
	//for i := 0; i < len(list1); i++ {
	//	list1[i] = i
	//	fmt.Println(list1[i])
	//}
	//for i := 0; i < len(list2); i++ {
	//	list2[i] = i
	//}
	//fmt.Println(list1)
	//fmt.Println(list2)
	//fmt.Println(len(list1))
	//fmt.Println(len(list2))
	//fmt.Println(cap(list2))
	//var slice1 = make([]int, 10)
	//fmt.Println(len(slice1))
	//var slise = make([]int, 10)
	//slicefib(slise)
	//bytes_test()
	//find_repleace()
	//make_str()
	//range_for()
	//fmt.Printf("%.2f", sum_slice(float_num))
	//resclice()
	////copy_append()
	//a := "卧槽尼玛"
	//flicestr := a[0:]
	//c := []byte(flicestr)
	//c[0] = 'c'
	//flicestr = string(c)
	//fmt.Println(flicestr)
	//sort_pkg()
	//h := "hello"
	//split_string(h, 4)
	//string_split2()
	//string_reverse()
	//Q29_uniq()
	//num := []int{0, 1, 2, 3, 4, 5, 6}
	//result := map_func(mult, num)
	//fmt.Println(result)
	//make_maps()
	//testelement()
	//map_days()
	//maps_forrange2()
	//sort_maps()
	//map_drink()
	//reboot_go()
	//fmt.Println(unsafe.Sizeof(100000))
	//package1.regjs()
	//morning := greetings.IsAM()
	//afternoon := greetings.Afternoon()
	//evening := greetings.Evening()
	//switch true {
	//case morning:
	//	greetings.GoodMorning()
	//case afternoon:
	//	greetings.GoodNight()
	//case evening:
	//	greetings.Evening()
	//}
	//for i := 0; i <= 100; i++ {
	//	if package1.Oddven(i) {
	//		fmt.Printf("%d是偶数\n", i)
	//	}
	//}
	structpractice.Main()
}
