package main

import "errors"
import "fmt"

//사용자 정의용 Error struct
type paramError struct {
	param int
	msg   string
}

//Error method 구현
func (e *paramError) Error() string {
	return fmt.Sprintf("%d - %s", e.param, e.msg)
}

//errors package를 이용한 error처리
func checkFunction1(param int) (int, error) {
	if param == 42 {
		return -1, errors.New("It is 42")
	}

	return param + 3, nil
}

//사용자 정의 paramError struct를 이용한 error처리
func checkFunction2(param int) (int, error) {
	if param == 42 {
		return -1, &paramError{param, "This param is error."}
	}

	return param + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := checkFunction1(i); e != nil {
			fmt.Println("checkFunction1 failed:", e)
		} else {
			fmt.Println("checkFunction1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := checkFunction2(i); e != nil {
			fmt.Println("checkFunction2 failed:", e)
		} else {
			fmt.Println("checkFunction2 worked:", r)
		}
	}

	_, e := checkFunction2(42)
	if ae, ok := e.(*paramError); ok {
		fmt.Println(ae.param)
		fmt.Println(ae.msg)
	}
}
