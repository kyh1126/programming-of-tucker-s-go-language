package main

import (
	"errors"
	"fmt"
)

//	func divide(a, b int) {
//		if b == 0 {
//			panic("b는 0일 수 없습니다") // Panic 발생
//		}
//		fmt.Printf("%d / %d = %d\n", a, b, a/b)
//	}
func divide(a, b int) error {
	if b == 0 {
		return errors.New("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	return nil
}

func main() {
	divide(9, 3)
	fmt.Print(divide(9, 0))
}
