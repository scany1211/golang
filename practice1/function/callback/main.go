package main

import (
	"fmt"

	even "github.com/self2/practice1/function/callback/example"
)

func Middle(k int, getNum func(int) int) int {
	fmt.Println("中间函数执行")
	return (1 + getNum(k))
}

func main() {
	k := 1
	i := Middle(k, even.Double)
	j := Middle(k, even.Triple)
	i2 := even.Double(k)
	j2 := even.Double(k)
	fmt.Println("起始函数：", i, j)
	fmt.Println("起始函数2：", i2, j2)
}
