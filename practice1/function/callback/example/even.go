// 回调函数

package even

import "fmt"

func Double(x int) int {
	fmt.Println("Double:", x)
	return x * 2
}

func Triple(x int) int {
	fmt.Println("Triple", x)

	return x * 3
}
