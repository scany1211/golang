package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
)

type Teacher struct {
	name    string
	age     int
	married bool
	sex     int
}

func main() {
	a := 10
	b := &a

	fmt.Printf("a has address is: %x \n", b)
	fmt.Printf("a has value is:  %d \n", a)
	*b++
	fmt.Printf("a has new value is: %d\n", a)

	c := Teacher{"steven", 32, true, 1}
	fmt.Printf("1. 变量c的内存地址是： %p. 值为： %v \n\n", &c, c)

	fmt.Printf("2. struct 变量的内存地址是： %p \n\n", c)

	changeStructVal(c)
	fmt.Printf("3. changeStructVal函数调用后的变量c的内存地址是： %p, 值为：%v \n\n", &c, c)
	changeStructPtr(&c)
	fmt.Printf("4. changeStructPtr函数调用后的变量c的内存地址是： %p, 值为：%v \n\n", &c, c)

}
func changeStructVal(c Teacher) {
	fmt.Printf("函数1内的c的地址是: %p, 值为：%v \n", &c, c)
	c.name = "test"
	c.age = 50
}
func changeStructPtr(c *Teacher) {
	fmt.Printf("函数2内的c的地址是: %p, 值为：%v \n", &c, c)
	(*c).name = "haha"
	(*c).age = 60
}

func ReadFileorStdin(filePath string) ([]byte, error) {
	var fileHandler *os.File

	if filePath == "-" {
		fileHandler = os.Stdin
	} else {
		var err error
		fileHandler, err = os.Open(filePath)
		defer fileHandler.Close()
		if err != nil {
			return nil, err
		}
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, fileHandler)

	return buf.Bytes(), nil
}
func validateConfig(configPath string) {
	log.Printf("Validating: %s\n", configPath)
	cmRaw, err := ReadFileorStdin(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var cm v1.ConfigMap

	err = yaml.UnmarshalStrict(cmRaw, &cm)
	if err != nil {
		log.Fatal(err)
	}

	config := NewConfig()
	if err := config.Load(&cm); err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration is valid.")
}
