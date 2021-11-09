package main

import (
	"fmt"
	"reflect"
)

func reflecttest(x interface{}) {
	//不知道调用函数会传入何种类型变量时，需要用空接口
	//1、通过类型断言
	// 这种方法只能去试

	// 2. 借助反射
	obj := reflect.TypeOf(x)
	fmt.Printf("type is %T, values is %v \n", obj, obj)
	fmt.Printf("type kind is %v, type name is %v \n", obj.Name(), obj.Kind())
}

//通过反射获取变量的值
func reflectvalue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("Valueof is %v, type is %T\n", v, v)
	k := v.Kind() // 值对应类型的准备
	switch k {
	case reflect.Int8:
		// 反射值转换为int32的变量
		ret := int8(v.Int())
		fmt.Printf("format is %v, %T\n", ret, ret)
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Printf("format is %v, %T\n", ret, ret)

	}
}

func reflectsetvalue(x interface{}) {
	v := reflect.ValueOf(x)
	// 用于根据指针取值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int8:
		v.Elem().SetInt(x:199 )
	case reflect.Float32:
		v.Elem().SetFloat( x: 3.1245 )
	}
}

type Cat struct{}

type Dog struct{}

func main() {
	var a float32 = 1.234
	var b int8 = 18
	reflecttest(a)
	reflecttest(b)

	// 结构体类型
	var c Cat
	var d Dog
	reflecttest(c)
	reflecttest(d)
	// slice 没有.Name
	var a1 []int
	var a2 []string
	reflecttest(a1)
	reflecttest(a2)
	fmt.Println("======================")
	reflectvalue(b)
	reflectvalue(a)
	fmt.Println("======================")
	reflectsetvalue(&a)
	reflectsetvalue(&b)

}
