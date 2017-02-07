package main

import "fmt"

const c int = 10
const s string = "Hi"

// const c = 10
// const s = "Hi"

const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "American Express"
)

//상수값을 0부터 순차적으로 부여하기 위해 iota 라는 identifier를 사용할 수 있다.
const (
	Apple  = iota // 0
	Grape         // 1
	Orange        // 2
)

func main() {
	var a string = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	//Short Assignment Statement ( := ) 라고 부르며 var를 생략하고 변수생성시 사용
	//함수 내부에서만 사용가능
	f := "short"
	fmt.Println(f)

	fmt.Println("Apple", Apple)
	fmt.Println("Grape", Grape)

	rawLiteral := `aaaaa\n 
    bbbb\n 
    cccc`

	interLiteral := "dddd\neeee"

	fmt.Println("rawLiteral", rawLiteral)
	fmt.Println("interLiteral", interLiteral)

	var iVal int = 100
	var uintVal uint = uint(iVal)
	var floatVal float32 = float32(iVal)
	println("=========", floatVal, uintVal)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println("===aa", bytes, str2)
}
