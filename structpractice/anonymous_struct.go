package structpractice

import "fmt"

type anonymousStruct struct {
	salary float32
	int
	string
}

func Anonymousestruct() {
	var anonstruct = anonymousStruct{
		salary: 1234.5,
		int:    9999,
		string: "工资",
	}
	fmt.Printf("这是匿名字段int:%d", anonstruct.int)
	fmt.Printf("这是匿名字段string:%s", anonstruct.string)
}
