package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func (f *File) Close() {
	fmt.Println("Close()")
}

//	func ReadFile(reader Reader) {
//		// Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환합니다.
//		// 런타임 에러가 발생합니다.
//		c := reader.(Closer)
//		c.Close()
//	}
func ReadFile(reader Reader) {
	// Reader 인터페이스 변수를 Closer 인터페이스로 타입 변환합니다.
	// 런타임 에러가 발생합니다.

	if c, ok := reader.(Closer); ok {
		c.Close()
	}
}

func main() {
	// File 포인터 인스턴스를 ReadFile() 함수의 인수로 사용합니다.
	file := &File{}
	ReadFile(file)
}
