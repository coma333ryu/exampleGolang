package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

//golang용 메서드 value receiver
//reat가 복사되어 전달된다.
func (r rect) area() int {
	r.width = 20
	return r.width * r.height
}

//golang용 메서드 포인터 receiver
//reat의 포인터가 전달된다.
func (r *rect) area2() int {
	r.width++
	return r.width * r.height
}

//struct는 mutable이고, 다른 함수의 파라메터로 전달시 Pass by Value에 따라 객체를 복사해서 전달하게 된다.
func changePersonVal(personStruct person) person {
	personStruct.name = "New Person"
	personStruct.age = 15
	return personStruct
}

//그래서 Pass by Reference로 struct를 전달하고자 한다면, struct의 포인터를 전달해야 한다.
func changePersonVal2(personStruct *person) *person {
	personStruct.name = "New Person"
	personStruct.age = 15
	return personStruct
}

func main() {
	//struct 생성
	firstPerson := person{}
	secondPerson := person{"Second Person", 11}
	//내장함수 new를 이용한 struct 생성
	//new를 이용하여 struct생성하면 반환되는 struct는 포인터값(*person)이 된다.
	thirdPerson := new(person)
	forthPerson := person{"Forth Person", 14}

	//person에 필드값 할당
	firstPerson.name = "First Person"
	firstPerson.age = 10

	thirdPerson.name = "Third Person"
	thirdPerson.age = 12

	fmt.Println("firstPerson", firstPerson)
	fmt.Println("secondPerson", secondPerson)
	fmt.Println("thirdPerson", *thirdPerson)

	newPerson := changePersonVal(forthPerson)
	fmt.Println("forthPerson 111====>", forthPerson)
	fmt.Println("newPerson", newPerson)

	newPerson2 := changePersonVal2(&forthPerson)
	fmt.Println("newPerson2", newPerson2)
	fmt.Println("forthPerson 222====>", forthPerson)

	println("========================= start method =========================")
	//reat struct가 복사되어 전달되기 때문에 원천 reat의 값에 영향을 주지 않는다.
	rect := rect{10, 20}
	area := rect.area()
	fmt.Println("rect width and area", rect.width, area)

	//reat 포인터가 전달되기 때문에 원천 reat가 변경되어진다.
	area2 := rect.area2()
	fmt.Println("rect width and area new ", rect.width, area2)
}
