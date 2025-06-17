package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) printpeople() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

func main() {
	fmt.Println("Method Example")

	person1 := Person{"Krishnendu", 20}
	person1.printpeople()
}
