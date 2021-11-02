package main

import (
	"fmt"
	"reflect"
)

func reflecttest(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("sy test type: %v, type value is: %v\n", v.Name(), v.Kind())
	k := reflect.ValueOf(x)
	fmt.Printf("sy test type: %v\n", k)

}

func main() {
	var f1 float32 = 3.14
	reflecttest(f1)

	var b int64 = 10
	reflecttest(b)

	type s1 struct {
		title   string
		publish int64
	}
	var s2 = s1{
		title:   "sy book infor",
		publish: 124,
	}
	reflecttest(s2) // type.Name is s1, type.kind is struct, valueof is "sy book info 123"

}
