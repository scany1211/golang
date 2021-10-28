package main

import (
	"fmt"
)

func main() {
	f1 := 1.23456 //默认float64
	fmt.Printf("%T\n", f1)
	f2 := float32(1.23456)
	fmt.Printf("%T, %v\n", f2, f2)
	b1 := true
	var b2 bool //默认为false
	c1 := &b2
	fmt.Printf("%T, %T, %#v \n", b1, b2, b2)
	fmt.Printf("%T\n%v\n", c1, c1)

}
