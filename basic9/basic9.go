package main

import "math"
import "fmt"

//Shape Interface 구현
//type명령어를 사용하고 interface라고 작성하면 됨.
// 메서드들의 집합
type Shape interface {
	area() float64
	perimeter() float64
}

//Rect 정의
type Rect struct {
	width, height float64
}

//Circle 정의
type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//Interface를 함수의 파라메터로 받을 수 있다.
func useShapeInterface(shape Shape) {
	fmt.Println("Single Shape", shape)
	fmt.Println("Single area method", shape.area())
	fmt.Println("Single perimeter method", shape.perimeter())
}

//멀티 파라메터도 가능함.
func useMultiShapeInterface(shapes ...Shape) {
	for _, shape := range shapes {
		areaMethod := shape.area()
		perimeterMethod := shape.perimeter()
		fmt.Println("Multi Shape", shape)
		fmt.Println("Multi area method", areaMethod)
		fmt.Println("Multi perimeter method", perimeterMethod)
	}
}

func printEmptyInterface(v interface{}) {
	fmt.Println(v)
}

func main() {
	rect := Rect{width: 3, height: 5}
	circle := Circle{radius: 4}

	useShapeInterface(rect)
	useShapeInterface(circle)

	useMultiShapeInterface(rect, circle)

	var x interface{}
	x = 1
	x = "Tom"

	printEmptyInterface(x)
}
