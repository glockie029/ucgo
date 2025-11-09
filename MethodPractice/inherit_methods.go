package MethodPractice

import "fmt"

type Base struct {
	id int
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float64
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(id int) {
	b.id = id
}

func (b *Base) GetId() {
	fmt.Printf("Id值为:%d", b.id)
}

func Inherit1() {
	e := new(Employee)
	e.SetId(1308)
	e.GetId()
}
