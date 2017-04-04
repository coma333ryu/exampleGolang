package tom_test

import "exampleGolang/oop/tom"
import "fmt"
import "testing"

func ExampleGetTotalMoney() {
	// godoc -http=:6060 실행 후 http://localhost:6060/pkg/exampleGolang/oop/tom/ 로 접속하여 보자.

	newTomStruct := tom.NewTom("TomExampler", 10, true)

	tomTotal := newTomStruct.GetTotalMoney(10, 200)
	fmt.Println(tomTotal)
	// Output: 20000
}

//go test를 이용하여 테스트가능
//특정 테스트 function만 실행하고 싶을때 go test -run 실행하고자 하는 function명 (ex: go test -v -run TestMove)
func TestMove(test *testing.T) {
	newTomStruct := tom.NewTom("TomTester", 10, true)
	newTomStruct.Move()
}
