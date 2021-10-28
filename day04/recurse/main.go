package main

import "fmt"

func curse(i int) (r1 int) {
	if i > 2 {
		r1 = curse(i-1) * i
	}
	fmt.Println("this is:", r1)
	return r1
}

func main() {
	result := 0
	for j := 2; j <= 10; j++ {
		result = curse(j)
		fmt.Println(result)
	}
}
