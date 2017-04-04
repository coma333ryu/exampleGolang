package tom_test

import (
	"exampleGolang/oop/tom"
	"fmt"
)

func ExampleGetTotalMoney() {
	// godoc -http=:6060 실행 후 http://localhost:6060/pkg/exampleGolang/oop/tom/ 로 접속하여 보자.

	newTomStruct := tom.NewTom("TomTester", 10, true)

	fmt.Println(newTomStruct.GetTotalMoney(10, 200))
	// Output: 20000
}
