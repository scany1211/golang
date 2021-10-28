package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	i2 := 5
	for ; i2 < 15; i2++ {
		fmt.Println(i2)
	}
	fmt.Println(i2)

	i3 := 3
	for i3 < 6 {
		fmt.Println(i3)
		i3++
	}
	fmt.Println(i3)

	// for range circle
	s := "hell沙河"
	for ss, v := range s {
		fmt.Printf("%d: %c\n", ss, v)
	}
	//输出9 9 乘法表
	for l1 := 1; l1 < 10; l1++ {
		for r1 := 1; r1 <= l1; r1++ {
			c1 := l1 * r1
			fmt.Printf("%v X %v=%v ", l1, r1, c1)
		}
		fmt.Printf("\n")
	}

}
