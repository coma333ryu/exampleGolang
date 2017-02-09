package main

import "fmt"

func main() {
	println("========================= start array =========================")
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[0])

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{1, 2, 3} //배열크기 자동으로

	println(a1[1])
	println(a3[2])

	var multiArr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	println("multiArr", multiArr[1][2])

	println("========================= start slice =========================")
	var slice1 []int        //슬라이스 변수 선언
	slice1 = []int{1, 2, 3} //슬라이스에 리터럴값 지정
	slice1[1] = 10
	fmt.Println(slice1)
	println(len(slice1), cap(slice1))

	//내장함수 make를 이용하여 슬라이스 생성
	//make 인자 ==> make(슬라이스 타입, 슬라이스 길이, 슬라이스 내부 배열의 최대 길이)
	slice2 := make([]int, 5, 10)
	println(len(slice2), cap(slice2)) // len 5, cap 10

	//슬라이스에 별도의 길이와 용량을 지정하지 않으면, 기본적으로 길이와 용량이 0 인 슬라이스를 만드는데, 이를 Nil Slice 라 하고, nil 과 비교하면 참을 리턴한다.
	var nilSlice []int

	if nilSlice == nil {
		println("Nil Slice")
	}
	println(len(nilSlice), cap(nilSlice))

	num := []int{1, 2, 3, 4, 5}
	fmt.Println("num =======> ", num)
	fmt.Println("len =======> ", len(num))
	fmt.Println("cap =======> ", cap(num))

	//append() ==> 슬라이스에 새로운 요소 추가.
	//append(슬라이스,추가할 값)
	//배열은 고정된 크기로 그 크기 이상의 데이타를 임의로 추가할 수 없지만, 슬라이스는 자유롭게 새로운 요소를 추가할 수 있다.
	//슬라이스 용량(capacity)이 아직 남아 있는 경우는 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이타를 추가하고,
	//용량(capacity)을 초과하는 경우 현재 용량의 2배에 해당하는 새로운 Underlying array 을 생성하고
	//기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당한다.
	num = append(num, 6)
	fmt.Println("new num =======> ", num)
	fmt.Println("new len =======> ", len(num))
	fmt.Println("new cap =======> ", cap(num))

	//슬라이스에 다른 슬라이스 추가하기.
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println("sliceA ======> ", sliceA)

	//슬라이스 복사
	sliceC := make([]int, len(sliceA))
	copy(sliceC, sliceA)
	fmt.Println("copy sliceC ====> ", sliceC)

	//부분 슬라이스(Sub-slice)
	//슬라이스[처음인덱스:마지막인덱스], 마지막인덱스는 원하는 인덱스+1
	//num[:3] ==> 첫 인덱스(0번째)부터 3번째까지 : 1,2,3 (마지막인덱스는 원하는 인덱스+1 을 사용)
	//처음 인덱스가 생략되면 0 이, 마지막 인덱스가 생략되면 그 슬라이스의 마지막 인덱스가 자동 대입된다.
	sliced1 := num[:3]
	fmt.Println("sliced1 =======> ", sliced1)
	fmt.Println("sliced1 len =======> ", len(sliced1))
	fmt.Println("sliced1 cap =======> ", cap(sliced1))

	sliced3 := sliced1[:4]
	fmt.Println("sliced3 =======> ", sliced3)
	fmt.Println("sliced3 len =======> ", len(sliced3))
	fmt.Println("sliced3 cap =======> ", cap(sliced3))
}
