package main

import "fmt"

func main() {
	names := []string{"bugs", "daffy", "elmer"}

	// Length
	fmt.Println(len(names))

	// First item (0 based indexing)
	fmt.Println(names[0])

	// Slicing
	fmt.Println(names[1:])

	// Traditional for loop
	for i := 0; i < len(names); i++ {
		fmt.Println(i)
	}

	// For loop - indices
	for i := range names {
		fmt.Println(i)
	}

	// For loop - indices & values
	for i, name := range names {
		fmt.Println(i, name)
	}

	// Append
	names = append(names, "tweety")
	fmt.Println(len(names))
}
