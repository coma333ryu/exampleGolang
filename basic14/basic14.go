package main

import "fmt"

func main() {
	// c := make(chan int)
	//수신루틴이 없으므로 데드락
	// c <- 1
	// fmt.Println(<-c)

	//버퍼채널 생성
	//make(chan type, N) : N는 생성할 버퍼의 갯수입력.
	ch := make(chan int, 1)

	//수신자가 없더라도 보낼 수 있다.
	ch <- 101

	fmt.Println(<-ch)
}
