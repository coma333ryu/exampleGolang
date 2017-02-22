package main

import "fmt"

//송신용 채널 pings설정
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//수신용 채널 pongs설정
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
