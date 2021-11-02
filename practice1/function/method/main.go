package main

import "fmt"

type tests struct {
	x, y float64
}

// tests 就是接受者，不能为内部类型，如int， float等
func (X tests) testm() float64 {
	return X.x + X.y
}

type tests2 struct {
	num int
}

// 值传递
func (X tests2) set1() int {
	X.num = 0
	return X.num
}

// 引用传递
func (X *tests2) set2() int {
	X.num = 2
	return X.num
}
func main() {
	t1 := tests{
		x: 3.14,
		y: 4.13,
	}

	fmt.Println(t1.testm())
	n1 := tests2{
		num: 100,
	}
	n1.set1()
	fmt.Println(n1.num)
	n1.set2()
	fmt.Println(n1.num)

	n2 := &tests2{num: 200}
	n2.set1()
	fmt.Println(n2.num)
	n2.set2()
	fmt.Println(n2.num)
}
