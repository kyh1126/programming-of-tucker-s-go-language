package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 크기 0인 채널 생성

	ch <- 9                    // main() 함수가 여기서 멈춘다
	fmt.Println("Never print") // 실행되지 않는다
}
