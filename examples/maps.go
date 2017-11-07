// Using maps
package main

import "fmt"

func main() {
	m := make(map[string]int)

	// Set values
	m["x"] = 1
	m["y"] = 2
	m["z"] = 3

	// Get value
	val := m["x"]
	fmt.Println(val)

	// Change value
	m["x"] = 2
	fmt.Println(m["x"])

	// Size
	size := len(m)
	fmt.Println(size)

	// Check if value exists
	_, ok := m["x"]
	if ok {
		fmt.Println("x in m")
	}

	// Delete values
	delete(m, "x")
	fmt.Println(len(m))

	// Iterate over keys
	for k := range m {
		fmt.Println(k)
	}

	// Iterate over keys & values
	for k, v := range m {
		fmt.Println(k, v)
	}

	// Create verbatim
	m1 := map[string]int{
		"x": 1,
		"y": 2,
		"z": 3,
	}
	fmt.Println(m1)
}
