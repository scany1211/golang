package main

import "fmt"

type inter1 interface {
	Impl1()
}

type struct1 struct {
	name string
}

func (s1 *struct1) Impl1() {
	if s1 == nil {
		fmt.Println("nil output")
		return
	}
	fmt.Println(s1.name)
}

func main() {
	var l1 inter1
	var ss1 *struct1
	l1 = ss1
	describe(l1)
	l1.Impl1()
	ss1 = &struct1{"hello world"}
	l1 = ss1
	describe(l1)
	l1.Impl1()
}

func describe(l1 inter1) {
	fmt.Printf("(%v, %T)\n", l1, l1)
}
