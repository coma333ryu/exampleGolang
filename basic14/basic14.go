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

	ch2 := make(chan int, 2)

	// 채널에 송신
	ch2 <- 1
	ch2 <- 2

	// 채널을 닫는다
	close(ch2)

	// 채널 수신
	println(<-ch2)
	println(<-ch2)

	//채널 수신시 return값은 채널 메세지와, 닫힘여부(true/false)이다.
	_, success := <-ch2

	if !success {
		println("더이상 데이타 없음.")
	}

}
