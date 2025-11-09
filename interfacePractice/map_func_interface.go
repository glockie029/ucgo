package interfacePractice

import "fmt"

type obj interface {
}

func mapFunc(obj obj) {
	fmt.Printf("开始处理数据,传入类型:%T\n", obj)
	switch v := obj.(type) {
	case int:
		doublev := v + v
		fmt.Printf("经过处理的%d的值为%d\n", v, doublev)
	case string:
		doublestring := v + v
		fmt.Printf("经过处理的%s的结果为%s\n", v, doublestring)
	}

}

func MapInterfaceMain() {
	IntAndString := []obj{
		123,
		"卧槽尼玛",
	}
	fmt.Println("将存在多个类型的数据传入mapFunc进行处理.....")
	for _, item := range IntAndString {
		mapFunc(item)
	}
}
