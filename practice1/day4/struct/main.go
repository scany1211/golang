package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type manager struct {
	person
	title string
}

func (p person) say() {
	fmt.Println("we will say hello to: ", p.name)
}

func main() {
	p1 := person{"James", 42}
	m1 := manager{p1, "Leader"}
	m1.say()
}
