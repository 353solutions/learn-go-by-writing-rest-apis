// for loops
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	n := 3
	for n > 0 {
		fmt.Println(n)
		n--
	}

	j := 0
	for {
		fmt.Println(j)
		j++
		if j > 2 {
			break // Stop the loop
		}
	}

	s := "abc"

	// Go over indices
	for i := range s {
		fmt.Println(i)
	}

	// Go over indices & values
	for i, c := range s {
		fmt.Println(i, c)
	}

}
