package main

import (
	"bytes"
	"fmt"
	"sort"
)

func find_repleace() {
	data := []byte("Hello world,hellp go!")
	if bytes.Contains(data, []byte("world")) {
		fmt.Println("数据中包括world")
	}
	newData := bytes.Replace(data, []byte("Hello"), []byte("Hi"), -1)
	fmt.Printf("原始数据:%s\n", data)
	fmt.Printf("转换后的数据:%s\n", newData)
}
func make_str() {
	//var buffer bytes.Buffer
	//buffer.WriteString("Go 语言")
	//buffer.WriteString("是一门")
	//buffer.WriteString("优")
	//buffer.WriteString("秀")
	//buffer.WriteString("的")
	//buffer.WriteString("的")
	//buffer.WriteString("语言")
	//
	//buffer.WriteByte(byte(' '))
	//buffer.WriteString("简单、高效、并发")
	//retult := buffer.String()
	//fmt.Println(retult)
	var b bytes.Buffer
	b.WriteString("hello")

	// 使用 Bytes()
	byteSlice := b.Bytes()
	fmt.Printf("通过 Bytes() 获取的内容：%s\n", byteSlice)

	// 这是一个陷阱！修改切片会影响 Buffer 的底层数据
	byteSlice[0] = 'H'
	fmt.Printf("修改后：通过 Bytes() 获取的内容：%s\n", byteSlice)
	fmt.Printf("Buffer 现在的状态：%s\n", b.Bytes()) // Buffer 的内容也被改变了

	fmt.Println("-----------------")

	// 重新初始化 Buffer
	b.Reset()
	b.WriteString("world")

	// 使用 String()
	str := b.String()
	fmt.Printf("通过 String() 获取的内容：%s\n", str)

	// 尝试修改字符串内容 (这会报错)
	// str[0] = 'W' // 编译错误：cannot assign to str[0]

	// 就算创建新的字符串，也不会影响 Buffer
	newStr := "W" + str[1:]
	fmt.Printf("新字符串：%s\n", newStr)
	fmt.Printf("Buffer 现在的状态：%s\n", b.Bytes()) // Buffer 保持不变
}

//func range_for() {
//	var slice []int = make([]int, 10)
//	for i := 0; i < 10; i++ {
//		slice[i] = i
//	}
//	for ix, value := range slice {
//		fmt.Printf("slice的第%d索引值是:%d\n", ix, value)
//	}
//	var slice1 []string = []string{"Hello", "world", "Go"}
//	for _, value := range slice1 {
//		fmt.Println(value)
//	}
//	//var slice3 [][]int = make([][]int, 10)
//	items := [...]int{10, 20, 30, 40, 50}
//	for i := 0; i < len(items); i++ {
//		fmt.Printf("第%d个元素\n", i)
//	}
//	for i := 0; i < 10; i++ {
//		items[i] = i
//	}
//	fmt.Println(items)
//}

var float_num []float32 = []float32{1.1, 2.2, 3.3}
var sum float32

func sum_slice(arrF []float32) float32 {
	for _, value := range arrF {
		sum += value
	}
	return sum
}
func resclice() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("切片的长度是:%d\n", len(slice1))
	}
	for n, value := range slice1 {
		fmt.Printf("slice %d:%d\n", n, value)
	}
}
func copy_append() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
func sort_pkg() {
	slice := []int{3, 6, 2, 4, 65, 1, 78}
	sort.Ints(slice)
	fmt.Println(slice)
}

// a = "Hello"
// {"H","e","l","l","o"}
// [:2]
func split_string(str string, i int) {
	str1 := str[0:i]
	str2 := str[i:]
	fmt.Println(str1, str2)
}
func string_split2() {
	str := "Google"
	str2 := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str, str2)
}
func string_reverse() {
	str := []byte("Google")
	str1 := make([]byte, len(str))
	for i := 0; i < len(str); i++ {
		target_index := i
		source_index := len(str) - 1 - i
		fmt.Println(source_index)
		str1[target_index] = str[source_index]
		fmt.Printf("str1的第%d:%s\n", i, str1)
	}
	fmt.Println(string(str1))
}
func Q29_uniq() {
	str := []string{"g", "o", "o", "g", "l", "e"}
	var str1 []string
	var lastval string
	for i, current := range str {
		if i == 0 {
			str1 = append(str1, current)
			lastval = current
			continue
		}
		if current != lastval {
			str1 = append(str1, current)
			lastval = current
		}
	}
	fmt.Println(str1)
	fmt.Println(len(str1))

}
func mult(num int) int {
	return num * 10
}
func map_func(mult func(int) int, num []int) []int {
	result := make([]int, len(num))
	for in, value := range num {
		result[in] = mult(value)
	}
	return result
}
