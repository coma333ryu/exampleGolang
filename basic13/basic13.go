package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		//<-를 이용하여 채널에 값 전달
		c <- "ping"
	}
}
func printer(c chan string) {
	for {
		//<-를 이용하여 채널의 값을 가져옴
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	//make를 이용하여 channel 생성
	c := make(chan string)

	//함수 pinger와 printer에 channel전달
	go pinger(c)
	go printer(c)

	//main함수도 goroutine이기 때문에 Scanln를 사용하여 goroutine이 종료되지 않게 막음.
	var input string
	fmt.Scanln(&input)

	/*	goroutine 실행 동기화를 위해 채널을 사용한 예
		done := make(chan bool)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
			}
			done <- true
		}()

		// 위의 Go루틴이 끝날 때까지 대기
		<-done
	*/
}
