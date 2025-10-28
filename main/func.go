package main

import (
	"fmt"
	"math"
	"strings"
)

const LIM = 41

var fibs [LIM]uint64

func name() {
	fmt.Println("hello world")
}

func test1() (a int, b int) {
	c := 1
	b = 20
	return c, b
}

func mult_returnval_01(a, b int) (int, int, int) {
	sum := a + b
	mult := a * b
	div := a / b
	return sum, mult, div
}

func mult_returnval_02(a, b int) (sum1 int, mult1 int, div1 int) {
	sum1 = a + b
	mult1 = a * b
	div1 = a / b
	return
}

func MySqrt1(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("%.02f不能为负数", x)
	}
	x1 := math.Sqrt(x)
	return x1, nil
}
func MySqrt2(x float64) (x2 float64, err error) {
	if x < 0 {
		return 0, fmt.Errorf("%.02f不能为负数", x)
	}
	x2 = math.Sqrt(x)
	return
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func newline(x ...int) {
	for _, i := range x {
		fmt.Println(i)
	}

}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func b() {
	fmt.Println("开始扫描...")
	defer fmt.Println("扫描结束正在统计数据..")
	fmt.Println("扫描中...")
	return
}
func fibonaccin(n int) (res1 int) {
	if n <= 1 {
		return 1
	}
	res1 = fibonaccin(n-1) + fibonaccin(n-2)
	return
}
func digui(num int) (n int) {
	if num <= 1 {
		return 1
	} else {
		fmt.Println(num)
		n = digui(num - 1)
	}
	return

}

func checkstring(s string) string {
	a := []rune(s)
	var resstr []rune
	for _, i := range a {
		if i <= 127 {
			resstr = append(resstr, rune(i))
		}
	}
	return string(resstr)
}

func printvalue() {
	sum, mult, div := mult_returnval_01(10, 20)
	sum1, mult1, div1 := mult_returnval_02(10, 20)
	fmt.Printf("非命名返回值sum: %d, mult: %d, div: %d\n", sum, mult, div)
	fmt.Printf("命名返回值sum1: %d, mult1: %d, div1: %d\n", sum1, mult1, div1)
	x1, err := MySqrt1(-12.3)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Printf("计算平方根(非命名参数):%.2f\n", x1)
	}
	x2, err := MySqrt2(-12.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("计算平方根(命名参数):%.2f\n", x2)
	}
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Printf("Multiply:%d\nn:%d\n", *reply, n)
	//================
	slice := []int{1, 2342, 2, 3, 4, 5}
	x := min(slice...)
	fmt.Println(x)
	newline(slice...)
	a()
	b()
	for i := 0; i <= 10; i++ {
		res := fibonaccin(i)
		fmt.Printf("第%d次的值为%d\n", i, res)
	}
	res1 := digui(10)
	fmt.Println(res1)
	s := "abcdefdhizZ9!@#$%^&*中文"
	runes := []rune(s)
	fmt.Println(runes)
	resstr := checkstring(s)
	fmt.Println(resstr)
}

func closepack() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf("- g is of type %T and has value %v\n", g, g)
	}
}
func test68() {
	fv := func() {
		fmt.Println("Hello,world")
	}
	fmt.Printf("fv type %T,the value is%v\n", fv)
}
func MakeAddscheme(prefix string) func(string) string {
	return func(url string) string {
		if !strings.HasPrefix(url, prefix) {
			return prefix + url
		}
		return url
	}
}
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
