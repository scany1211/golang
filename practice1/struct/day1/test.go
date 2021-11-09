//描述某些属性， 把基本数据类型组合在一起，类似于class

package main

import "fmt"

// 定义结构体
// type person struct {
// 	name string
// 	age int
// 	marriage bool
// 	city string
// }
// 简写，结构体内存对齐。
type person struct {
	name, city string
	age        int
	marriage   bool
}

// 结构体的实例化， 实例化后才会分配内存。
func main() {
	var p1 person
	p1.name = "david"
	p1.city = "beijing"
	p1.age = 18
	p1.marriage = false

	fmt.Printf("values is %v\n", p1)
	fmt.Println(p1.name, p1.city)
	//匿名结构体， 用于符合类型，临时场景，结构体只使用一次
	var user struct {
		name     string
		marriage bool
	}
	user.name = "dfett"
	user.marriage = true

	fmt.Printf("values is %v\n", user)
}
