package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

type OpFn func(int, int) int

func getOperator(op string) OpFn { // op에 따른 함수 타입 반환
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator OpFn // int 타입 인수 2개를 받아서 int 타입 반환을 하는 함수 타입 변수
	operator = getOperator("*")

	var result = operator(3, 4) // 함수 타입 변수를 사용해서 함수 호출
	fmt.Println(result)
}
