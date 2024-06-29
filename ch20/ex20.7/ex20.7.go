package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

func PrintAge(stringer Stringer) {

	s := stringer.(*Student) // *Student 타입으로 타입 변환
	fmt.Printf("Age: %d\n", s.Age)

}

func main() {
	s := &Student{15} // *Student 타입 변수 s 선언 및 초기화

	PrintAge(s) // 변수 s 를 인터페이스 인수로 PrintAge() 함수 호출
}
