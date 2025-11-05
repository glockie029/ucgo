package structpractice

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func Main() {
	var pers1 Person
	pers1.firstName = "Chris"
	pers1.lastName = "Wooward"
	upPerson(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	per2 := new(Person)
	per2.firstName = "Chris"
	per2.lastName = "Wooward"
	(*per2).lastName = "Woodward"
	upPerson(*per2)
	fmt.Printf("The name of the person is %s %s\n", per2.firstName, per2.lastName)

	per3 := &Person{"Chirs", "Woodward"}
	upPerson(*per3)
	fmt.Printf("The name of the person is %s %s\n", per3.firstName, per3.lastName)
}
