package main

//字符串是用双引号表示，单引号表示的是字符
//字符： 比如 c := 'h'，'啥'，一个汉字也是一个字符；一个字符=1个字节，1一个utf8编码的汉字，占一般三个字节
//一个字节=8个Bit位
//转义符号反斜杠\
import (
	"fmt"
	"strings"
)

func main() {
	path1 := "\"d:\\teststaudy\\day1"
	fmt.Println(path1)
	fmt.Printf("%v\n", path1)

	//多行字符串,反引号
	s1 := `
test1
efcho2
"d:\teststaudy\day1
`
	fmt.Println(s1)
	// 拼接字符串
	fmt.Println(len(s1))

	fmt.Println(s1 + path1)
	s2 := fmt.Sprintf("%s%s", s1, path1)
	fmt.Println(s2)

	//切割
	s3 := strings.Split(path1, "\\")
	fmt.Printf("%T\n%v\n", s3, s3)

	fmt.Println(strings.HasPrefix(path1, "\""))

	fmt.Println(strings.Join(s3, "+"))

	//字符串修改只能转换为其它类型修改

	s4 := "heello"
	s5 := []rune(s4) //把字符串强制转换为一个切片, 切片元素为字符。rune本质上是init32类型
	fmt.Println(string(s5))
	s5[0] = 't'
	fmt.Println(string(s5))
	c1 := "H"
	c2 := 'H'
	fmt.Printf("c1:%T c2:%T\nc2 value: %d\n", c1, c2, c2)

}
