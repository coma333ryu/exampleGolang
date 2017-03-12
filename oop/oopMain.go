package main

import (
	"exampleGolang/oop/human"
	"exampleGolang/oop/tom"
	"fmt"
)

func main() {
	human := human.NewHuman("Human", 10, true)
	fmt.Println(human)
	human.Move()
	human.Communicate()

	tom := tom.NewTom("Tom", 20, true)
	fmt.Println("Tom", tom)
	fmt.Println("Tom's Total Money is ", tom.GetTotalMoney(10, 100))
	tom.Move()
	tom.Communicate()
}
