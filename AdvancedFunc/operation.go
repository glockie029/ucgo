package AdvancedFunc

import "fmt"

// 函数类型别名,接受两个证书并返回一个证书的函数签名
type Operation func(int, int) int

func DoOperation(a int, b int, op Operation) int {
	return op(a, b)
}
func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func OprattiongMain() {
	result := DoOperation(1, 2, Add)
	fmt.Println(result)

}
