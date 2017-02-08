package main

//func : 함수 키워드
//Pass By Value 파라메터 전달방식
func say(msg string) {
	println(msg)
}

//Pass By Reference
//*string 과 같이 파라미터가 포인터임을 표시
//say2함수 내부의 msg는 문자열이 아니라 문자열을 갖는 메모리 영역의 주소값
func say2(msg *string) {
	println(*msg)
	*msg = "World" //메시지 변경 (Dereferencing)
}

// Variadic Function (가변인자함수)
func say3(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

//함수에 리턴값이 있을 경우 리턴 타임을 파라메터() 다음에 기술한다.
//함수의 마지막에 return문과 함께 리턴값을 기술한다.
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sum2(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

//Named Return Parameter : 리턴값들을 할당하여 리턴
//count와 total에 함수내부에서 직접 값을 할당함.
//return 문에는 아무 값들을 리턴하지 않지만, 그래도 리턴되는 값이 있을 경우에는 빈 return 문을 반드시 써 주어야 한다 (이를 생략하면 에러 발생).
func sum3(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

//재귀함수
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	msg := "Hello"
	say(msg)
	//변수앞에 & 부호를 붙이면 msg 변수의 주소를 표시
	//msg 변수의 주소를 전달
	say2(&msg)

	println("======", msg) //변경된 메시지 출력

	say3("This", "is", "a", "book")
	say3("Hi")

	total := sum(1, 7, 3, 5, 9)
	println("total ========>", total)

	count, total := sum2(1, 7, 3, 5, 9)
	println("count and total ======> ", count, total)

	count1, total1 := sum3(1, 7, 3, 5, 9)
	println("count1 and total1 ======> ", count1, total1)

	println("fact(7)", fact(7))
}
