// goroutines
package main

import (
	"fmt"
	"time"
)

func Printer(name string) {
	fmt.Println("Hi, I'm", name)
}

func main() {
	names := []string{"Bugs", "Daffy", "Elmer"}

	for _, name := range names {
		go Printer(name)
	}

	// Wait for the goroutines to finish (there are better ways, see sync.WaitGroup
	time.Sleep(time.Millisecond)
}
