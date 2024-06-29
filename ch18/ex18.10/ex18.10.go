package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2 // 삭제할 인덱스

	copy(slice[idx:], slice[idx+1:])

	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
