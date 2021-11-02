package main

import (
	"errors"
	"fmt"
)

type stype1 struct {
	x, y int
}

type iinterf interface {
	mmthod1() int
}

func (s stype1) mmthod1() int {
	return s.x + s.y
}

type Cari interface {
	drive(Driver)
}

type Driver struct {
	name string
}

type Benz struct{}
type Audi struct{}

func (this *Benz) drive(d Driver) {
	fmt.Println("this is benz driver: ", d.name)
}

func (this *Audi) drive(dr Driver) {
	fmt.Printf("%s drives Audi.\n", dr.name)
}

func Newcar(c string) (Cari, error) {
	switch c {
	case "Benz":
		return &Benz{}, nil
	case "Audi":
		return &Audi{}, nil
	default:
		return nil, errors.New("not support")
	}

}

func main() {
	var s1 iinterf = &stype1{x: 3, y: 4}
	var s2 iinterf
	s3 := stype1{x: 3, y: 4}
	s2 = &s3
	fmt.Println(s1.mmthod1())
	fmt.Println(s2.mmthod1())

	mycar := Driver{name: "davide"}
	car, error := Newcar("Benz")
	if error != nil {
		fmt.Println(error)
	} else {
		car.drive(mycar)
	}
}
