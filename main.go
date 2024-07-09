package main

import (
	"fmt"
	"github.com/Kay7271/demo/mypackage"
)

var b = 3

const (
	pi = 3.1415926
	e  = 2.71828
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	_ = iota * iota
	a
	c
	d
)

func main() {

	name := "Kay"

	var age int
	age = 12

	var gender string = "female"

	var id = 1

	mypackage.New()
	fmt.Printf("ID: %d,", id)
	fmt.Printf("Name: %s,", name)
	fmt.Printf("Age: %d,", age)
	fmt.Printf("Gender: %s", gender)

	a := 1
	_, sub := ope(a, b)
	fmt.Printf("Sub: %d", sub)

	fmt.Println(d)

}

func ope(a, b int) (sum, sub int) {
	return a + b, a - b
}
