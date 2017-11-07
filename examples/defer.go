// defer examples
package main

import (
	"fmt"
	"os"
)

func cleanup() {
	fmt.Println("cleanup")
}

func caller() {
	defer cleanup()

	fmt.Println("caller code")
}

func main() {
	caller()

	file, err := os.Open("defer.go")
	if err != nil {
		fmt.Printf("error: can't open - %s\n", err)
		return
	}
	defer file.Close()

	// Work with file here

}
