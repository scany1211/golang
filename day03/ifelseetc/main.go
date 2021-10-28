package main

// if条件判断
import (
	"fmt"
)

func main() {
	age := 20
	if age > 18 {
		fmt.Println("确实长大了")
	} else {
		fmt.Println("确实meiyou长大了")

	}

	// 多个条件判断
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年了")

	} else {
		fmt.Println("还小哈")

	}

	//特殊写法，一下的作用变量只在if中生效
	if age2 := 19; age2 > 18 {
		fmt.Println("确实长大了")
	} else {
		fmt.Println("还小哈")
	}
	// 程序外找不到age2

}
