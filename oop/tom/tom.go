package tom

import (
	"exampleGolang/oop/human"
	"fmt"
)

//Tom : Human struct를 상속받음.
type Tom struct {
	humanTom    human.Human
	workingTime int
	money       int
	total       int
}

//NewTom : Tom struct 생성자
func NewTom(name string, age int, sex bool) *Tom {
	return &Tom{
		humanTom: *human.NewHuman(name, age, sex),
	}
}

func (tom *Tom) work(time int, money int) {
	tom.total = time * money
	fmt.Println("Tom is working")

}

//Move : Humna의 Move 메소드 오버로딩
func (tom *Tom) Move() {
	tom.humanTom.Move()
	fmt.Println("Tom is moving!!!")
}

//Communicate : Humna의 Communicate 메소드 오버로딩
func (tom *Tom) Communicate() {
	tom.humanTom.Communicate()
	fmt.Println("Tom is communicating!!!")
}

//GetTotalMoney :
func (tom *Tom) GetTotalMoney(time int, money int) int {
	tom.work(time, money)
	return tom.total
}
