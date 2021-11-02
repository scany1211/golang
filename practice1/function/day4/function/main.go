package main

import "fmt"

// main没有返回值
// 没有默认参数

// ret 变量可以直接使用, 命名返回值就相当于在函数中已经声明一个变量， 如果使用命名返回值，return可以不写
func sum(x int, y int) (ret int) {
	ret = x + y
	return
}
func f1(x int, y int) {
	fmt.Println("this is function1")
}

func f2() {
	fmt.Println("function f2")
}

// 没有命令返回值，需要显示return
func f3() int {
	return 2
}

//多个返回时
func f4() (int, string) {
	return 1, "yesy"
}

//参数类型一致时，可以省略前面类型的参数

func f5(x, y, z int) int {
	return x + y + z
}

// y的类型为切片, 可变长参数，必须放在函数参数的最后
func f6(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	r := sum(10, 20)
	f1(1, 2)
	f2()
	f3()
	n, _ := f4()

	f6("test", 1, 2, 3, 4, 5)
	fmt.Println("this is f4", n)
	fmt.Println(r)
}
