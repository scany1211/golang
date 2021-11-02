package main

import "fmt"

type testI interface {
	add()
}

type testS struct {
	x, y int
}

func (t testS) add() {
	fmt.Println(t.x + t.y)
}

func main() {
	var a1 interface{} = 100
	v, ok := a1.(int)
	fmt.Println(v, ok) // v的值是100，ok为true
	fmt.Printf("%T\n", v)

	v2, ok := a1.(float64)
	fmt.Println(v2, ok) // v2 为0， ok为false， 因为floated64的初始值为0.
	fmt.Printf("%T\n", v2)

	s1 := testS{x: 3, y: 4}
	var e interface{}
	e = s1
	fmt.Println("========================")
	fmt.Printf("e has type: %T value: %v\n", e, e)
	fmt.Println("========================")

	x, ok := e.(testS)
	if ok == true {
		fmt.Printf("x has type: %T, value: %v\n", x, x)
		fmt.Println("========================")
		x.add()
	}
	var l2 testI
	l2 = x
	fmt.Printf("l2 has type: %T, value: %v\n", l2, l2)
	l2.add()
	fmt.Println("========================")

	var l3 testI
	l3 = e.(testI)
	fmt.Printf("l3 has type: %T, value: %v\n", l3, l3)
	l3.add()
	fmt.Println("========================")

	var l4 testI
	l4 = l2.(testI)
	fmt.Printf("l4 has type: %T, value: %v\n", l4, l4)
	l3.add()
	fmt.Println("========================")

}
