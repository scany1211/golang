package main

import "fmt"

//定义接口的时候并没有定义其实现
type error interface {
	errorMethod() string
}

type errorMsg struct {
	code    int
	message string
}

type duck interface {
	walk()
	quak()
}

type cat struct{}

// func (c cat) walk() {
// 	fmt.Println("strcut cat walk")
// }

// func (c cat) quak() {
// 	fmt.Println("strcut cat quak")
// }

func (c *cat) walk() {
	fmt.Println("strcut cat walk pointer")
}

func (c *cat) quak() {
	fmt.Println("strcut cat quak pointer")
}

func main() {
	message := errorMsg{400, "unknow error"}
	out := message.errorMethod() //调用接口
	fmt.Println("output is ", out)
	type test struct{}
	v := test{}
	fmt.Printf("Before type and value is: %t, %v \n", v, v)
	printTest(v)
	var c1 duck = &cat{} // 使用结构体指针初始化变量
	var c2 duck = cat{}  // 使用结构体初始化变量
	c1.walk()
	c2.walk()
}

// 具体接口的实现
func (m *errorMsg) errorMethod() string {
	return m.message
}

func printTest(v interface{}) {
	fmt.Printf("empty interface test: %v, %t\n", v, v)
}
