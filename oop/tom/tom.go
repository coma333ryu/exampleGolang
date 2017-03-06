package tom

import (
	"exampleGolang/oop/human"
	"fmt"
)

type Tom struct {
	humanTom    human.Human
	workingTime int
	money       int
	total       int
}

func NewTom(name string, age int, sex bool) Tom {
	return Tom{
		humanTom: human.NewHuman(name, age, sex),
	}
}

func (tom *Tom) work(time int, money int) {
	tom.total = time * money
	tom.humanTom.Move("Tom")
	fmt.Println("Tom is working")

}

func (tom *Tom) GetTotalMoney(time int, money int) int {
	tom.work(time, money)
	return tom.total
}
