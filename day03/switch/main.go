package main

import "fmt"

func main() {
	// i=5是跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over for circle")

	// i=5是跳出本次for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //跳出for循环
		}
		fmt.Println(i)
	}
	fmt.Println("over for this circle")
	// n := 2
	// sySwitch(n)
	syGoto()

}

func sySwitch(a int) {
	switch a {
	case 1:
		fmt.Println("this is num 1")
	case 2:
		fmt.Println("this is num 2")
	default:
		fmt.Println("this is default")

	}
}

//goto 跳出多层循环

func syGoto() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto lb1 //调到指定标签处
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
lb1:
	fmt.Println("goto here")
}
