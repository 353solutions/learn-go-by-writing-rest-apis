package main

import "fmt"

func main() {
	i := 4

	if i > 0 {
		fmt.Println("i is big")
	}

	if i > 0 && i < 10 {
		fmt.Println("i in range")
	}

	if i != 5 {
		fmt.Println("i is not 5")
	}
}
