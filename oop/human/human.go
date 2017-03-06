package human

import "fmt"

//Human struct
type Human struct {
	name string
	age  int
	sex  bool
}

//NewHuman : Human struct 생성자
func NewHuman(name string, age int, sex bool) Human {
	return Human{name: name, age: age, sex: sex}
}

//Move method 구현
func (human *Human) Move(humanName string) {
	fmt.Printf("%q is moving \n", humanName)
}

//Communicate method 구현
func (human *Human) Communicate() {
	fmt.Println("Human is communicating")
}
