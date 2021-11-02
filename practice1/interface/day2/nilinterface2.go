package main

import "fmt"

type Inter1 interface {
	Impl()
}

type Stru []int

func (s Stru) Impl() {
	fmt.Printf("%T\n", s)
	if s == nil {
		fmt.Println("<nil>")
		return
	}
}

func main() {
	var s1 Stru
	fmt.Printf("%T, %v\n", s1, s1)
	if s1 == nil {
		fmt.Println("s1 is nil")
	}

	var s2 Inter1
	if s2 == nil {
		fmt.Println("s2 is nil")
	}

	s2 = s1
	if s2 == nil {
		fmt.Println("s1 is nil")
	}

}
