package interfacePractice

import (
	"fmt"
	"reflect"
)

func ReflectMain() {
	var x float64 = 3.4
	fmt.Println("ValueOf方法====>type:", reflect.ValueOf(x).Float())
	v := reflect.ValueOf(&x)
	v = v.Elem()
	v.SetFloat(3.1415)
	fmt.Println(v.CanSet())
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println("Interface:", v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
