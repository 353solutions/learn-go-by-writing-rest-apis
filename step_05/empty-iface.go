package main

import "fmt"

func main() {
	var i interface{}

	i = 10
	// Print value and type
	fmt.Printf("%v (%T)\n", i, i) // 10 (int)

	i = "hello"
	fmt.Printf("%v (%T)\n", i, i) // hello (string)
}
