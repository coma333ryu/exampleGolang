package main

import (
	"fmt"
	"time"
)

func main() {
	println("========================= start if =========================")
	//()가 없어도 된다.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//if문 내부에서 간단한 실행문을 사용할 수 있음.
	//num값의 scope는 해당 if문 내에서만 사용가능.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// println("num======>", num)

	println("========================= start switch =========================")
	//switch
	//case문에 break문이 필요없음. go컴파일러가 자동으로 break 문을 각 case문 블럭 마지막에 추가하시 때문임.
	//fallthrough문을 작성하여 아래의 casse문들을 실행 할 수 있음.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("three")
	}

	//case문에 여러개의 표현식을 사용할 수 있음.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//case문에 표현식을 작성 할 수 있음.
	score := 100
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}

	//switch문에 표현식이 없을 수도 있음.
	//이러한 경우 switch가 true라 생각하고 첫번째 case문을 실행함.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만, go는 그 변수의 Type에 따라 case로 분기할 수 있다.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	println("========================= start for =========================")
	//golang에서의 반복문은 for만 존재함.
	//()는 사용하지 않음. 사용시 컴파일 에러발생함.
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println("sum =========>", sum)

	//조건식만 사용하여도 가능함. (타 언어의 while문과 같이 사용하기 위해)
	num := 1
	for num <= 3 {
		fmt.Println(num)
		num = num + 1
	}

	//무한루프
	// for {
	//     println("Infinite loop")
	// }

	//for range문 (foreach와 비슷)
	names := []string{"AAA", "BBB", "CCC"}

	for index, name := range names {
		println(index, name)
	}
	//index가 필요 없을시 _(공백 식별자)를 사용하여 무시할 수 있음.
	for _, name := range names {
		println(name)
	}

	//break : for문을 빠져나올 때 사용
	//continue : for 루프의 중간에서 나머지 문장들을 실행하지 않고 for 루프 시작부분으로 가려할 때 사용.
	//goto : 기타 임의의 문장으로 이동하기 위해 사용.
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue // for루프 시작으로
		}
		a++
		if a > 10 {
			break //루프 빠져나옴
		}
	}
	if a == 11 {
		goto END //goto 사용예
	}
	println(a)

END:
	println("End")

	//break문이 레이블과 같이 사용되어 특정 레이블로 이동이 가능.
	idx := 0
L1:
	for {
		if idx == 0 {
			break L1
		}
	}

	println("OK")
}
