package main

import "fmt"

//空接口可以赋任何值
type I interface{}

func main() {
	var l1 I
	l1 = 10
	Print(l1)
	l1 = "hello"
	Print(l1)

	slic1 := []int{1, 2, 3, 4, 5, 6, 7}
	Print(slic1) // type: []int, value: [1 2 3 4 5 6 7]
	l1 = slic1
	Print(l1) // type: []int, value: [1 2 3 4 5 6 7]， 注意l1不能再做切片了

	f1 := slic1[0:3]
	Print(f1) // type: []int. value: [1,2,3]
	f1[1] = 50
	Print(f1) //type: []int, value: [1 50 3]
}

func Print(inter1 I) {
	fmt.Printf("this is type %T \n value is %v\n", inter1, inter1)
}
