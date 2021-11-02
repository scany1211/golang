package main

import "fmt"

type s1 struct{}

func nilOrNot(v interface{}) bool {
	return v == nil
}

func main() {
	var s *s1                                  // 初始化后是nil
	fmt.Println("s1 is nil or not:", s == nil) //初始化了一个 *TestStruct 类型的变量，由于指针的零值是 nil，所以变量 s 在初始化之后也是 nil：
	fmt.Println(nilOrNot(s))                   //*TestStruct 类型会转换成 interface{} 类型，转换后的变量不仅包含转换前的变量，还包含变量的类型信息 TestStruct，所以转换后的变量与 nil 不相等。
}
