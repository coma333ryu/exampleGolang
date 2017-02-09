package main

import (
	"fmt"
)

func main() {
	//map를 생성하기 위해서는 make()함수를 사용해야 한다.
	var idMap = make(map[int]string)

	idMap[0] = "Zero"
	idMap[1] = "One"
	idMap[2] = "Two"

	fmt.Println("idMap", idMap)

	//리터럴(literal)을 사용한 초기화
	idMap2 := map[int]string{
		0: "Zero Val",
		1: "One Val",
		2: "Two Val",
	}
	fmt.Println("idMap2", idMap2)

	//map 데이터 읽기
	fmt.Println("idMap[0]", idMap[0])
	fmt.Println("idMap2[1]", idMap2[1])

	//map 삭제
	//delete(삭제할 map,삭제할 key value)
	delete(idMap, 1)

	fmt.Println("after delete idMap", idMap)

	//map key의 존재유무
	// val : 해당 map의 value, exists : 해당 key의 존재유무(true,false)
	// val, exists := idMap2[3]
	_, exists := idMap2[3]
	if !exists {
		println("No idMap2[3] key")
	}

	//for range를 이용한 map데이터 읽기
	for key, val := range idMap2 {
		fmt.Println("idMap2 range", key, val)
	}
}
