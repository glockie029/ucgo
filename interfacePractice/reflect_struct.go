package interfacePractice

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func Reflectstruct() {
	value := reflect.ValueOf(secret)
	fmt.Println(value)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)
	fmt.Println(value.NumField())
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("field %d:%s\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
