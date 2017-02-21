package main

import (
	"fmt"
	"os"
)

func firstStart() {
	fmt.Println("First Start Function")
}

func secondStart() {
	fmt.Println("Second Start Function")
}

func thirdStart() {
	defer fmt.Println("Third End Function")
	fmt.Println("Third Start Function")
}

func openFile(fn string) {
	//panic 함수에 의한 패닉상태를 다시 정상상태로 되돌리는 함수
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	//현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴
	if err != nil {
		panic(err)
	}

	defer f.Close()

}

func main() {
	thirdStart()
	//특정 문장 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행
	defer secondStart()
	firstStart()

	openFile("Invalid.txt")
	fmt.Println("Done")
}
