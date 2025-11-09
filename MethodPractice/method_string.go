package MethodPractice

import (
	"fmt"
	"strconv"
)

type Kobe struct {
	Name string
	Age  int
}

func MethodStringMain() {
	p := new(Kobe)
	p.Name = "KobeBryant"
	p.Age = 42
	fmt.Printf("科比:%v", p)
}

func (p *Kobe) String() string {
	return "(科比的英文名:" + p.Name + "科比的年龄：" + strconv.Itoa(p.Age) + ")"
}
