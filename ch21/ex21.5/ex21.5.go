package main

import "fmt"

func main() {
	i := 0

	f := func() {
		i += 10
	}

	i++

	f() // f 함수 타입 변수를 사용해서 함수 리터럴 실행

	fmt.Println(i)
}
