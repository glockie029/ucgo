package structpractice

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func Structpractice() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	fmt.Printf("This int is:%d\n", ms.i1)
	fmt.Printf("This string is:%f\n", ms.f1)
	fmt.Printf("This string is:%s\n", ms.str)
	fmt.Println(ms)
	ma := &struct1{1, 15.5, "Chris"}
	fmt.Println(ma)
}
