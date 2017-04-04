/*
package tom 은 human을 상속받은 tom에 관한 struct들과 메소드들의 패키지임.
*/
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

//Move : Humna의 Move 메소드 오버라이딩
func (tom *Tom) Move() {
	tom.humanTom.Move()
	fmt.Println("Tom is moving!!!")
}

//Communicate : Humna의 Communicate 메소드 오버라이딩
func (tom *Tom) Communicate() {
	tom.humanTom.Communicate()
	fmt.Println("Tom is communicating!!!")
}

//GetTotalMoney :
func (tom *Tom) GetTotalMoney(time int, money int) int {
	tom.total = time * money
	return tom.total
}
