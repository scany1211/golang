package main

import "fmt"

func typetest(a int, b string, c []int, d [2]int, e map[int]string, f func(int, int), g chan int, h interface{}) {
	fmt.Printf("address is ： a=%p, b=%p, c=%p, d=%p, e=%p, f=%p, g=%p,h=%p\n", &a, &b, &c)
}
func main() {
	a := 10
	b := "holdhold0"
	c := []int{1, 2, 3, 4, 5}
	d := [2]int{2, 3}
	e := make(map[int]string, 3)
	f := func(x, y int) {}
	g := make(chan int, 10)
	var h interface{}
	fmt.Printf("defined values is: a=%p, b=%p, c=%p, d=%p, e=%p, f=%p, g=%p,h=%p\n", &a, &b, &c, &d, &e, &f, &g, &h)
	typetest(a, b, c, d, e, f, g, h)

}
