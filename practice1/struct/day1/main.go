package main

import "fmt"

type student struct {
	name string
	age  int64
}

// teacher会继承student的属性
type teacher struct {
	student
	classname string
}

func (s student) sayhello() {
	fmt.Println("student name: ; age: ", s.name, s.age)
}

func (t teacher) sayhello() {
	fmt.Println("teacher name: ; age: ", t.name, t.age)
}
func main() {
	s1 := student{
		name: "daniel",
		age:  14,
	}
	t1 := teacher{}
	t1.name = "davide"
	t1.age = 44
	t1.classname = "test"

	s1.sayhello()
	t1.sayhello()

}
