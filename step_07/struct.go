package main

import "fmt"

type Loon struct {
	Name string
	Age  int
}

func NewLoon(name string, age int) *Loon {
	return &Loon{name, age}
}

func main() {
	elmer := Loon{"Elmer", 80}
	fmt.Println(elmer)

	daffy := Loon{
		Age:  80,
		Name: "Daffy",
	}
	fmt.Println(daffy)

	fmt.Println(NewLoon("Tweety", 75))
}
