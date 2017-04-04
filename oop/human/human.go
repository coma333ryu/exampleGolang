package human

import "fmt"

//Human struct
type Human struct {
	name string
	age  int
	sex  bool //true : men, false : women
}

//NewHuman : Human struct 생성자
func NewHuman(name string, age int, sex bool) *Human {
	return &Human{name: name, age: age, sex: sex}
}

//Move method
func (human *Human) Move() {
	fmt.Printf("%q is moving \n", human.name)
}

//Communicate method
func (human *Human) Communicate() {
	fmt.Printf("%q is communicating \n", human.name)
}
