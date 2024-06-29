package main

import "fmt"

func main() {
	m := make(map[int]int) // 맵 생성
	m[1] = 0               // 요소 추가
	m[2] = 2
	m[3] = 3

	delete(m, 3) // 요소 삭제
	delete(m, 4) // 없는 요소 삭제 시도

	v, ok := m[3]
	fmt.Println(v, ok) // 삭제된 요소 값 출력
	v, ok = m[1]
	fmt.Println(v, ok) // 존재하는 요소 값 출력
}
