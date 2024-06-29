package main

import "fmt"

type myInt int

func (m myInt) Add(a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func addFunc(m myInt, a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func main() {
	var a myInt
	fmt.Println(a.Add(10))
	fmt.Println(addFunc(a, 10))
}
