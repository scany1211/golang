package main

import "fmt"

//切片的本质就是一个框，框住了一块连续的内存。只能保存相同类型的变量
func main() {
	// m1 = [...]int
	// var ss1 []int //len=0, cap=0, ss1=nil
	// ss2 :=[]int //len=0, cap=0, ss1 !=nil
	// ss3 := make([]int, 0) //len=0, cap=0, ss1 !=nil
	// s1 := make([]int, 5, 10)
	// fmt.Printf("s1=%v s1.len: %d s1:cap: %d", s1, len(s1), cap(s1))
	// fmt.Printf("ss1=%v s1.len: %d s1:cap: %d", ss1, len(ss1), cap(ss1))
	// fmt.Printf("ss2=%v s1.len: %d s1:cap: %d", ss2, len(ss2), cap(ss2))
	// fmt.Printf("ss3=%v s1.len: %d s1:cap: %d", ss3, len(ss3), cap(ss3))
	s2 := []int{1, 3, 5}
	s3 := s2
	fmt.Println(s2, s3)
	s3[0] = 1000
	fmt.Println(s2, s3)
	// 切片的便利
	//1. 索引
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}

	for i, v := range s3 {
		fmt.Println(i, v)
	}
	//调用append函数必须用原来的切片变量接收返回值
	s3 = append(s3, 300)
	fmt.Println(s3)

}
