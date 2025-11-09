package MethodPractice

import "fmt"

type Basec struct{}

func (Basec) Magic() {
	fmt.Println("base magic")
}
func (self Basec) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Basec
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
func MagicMain() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
