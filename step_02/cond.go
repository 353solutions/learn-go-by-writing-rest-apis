package main

import "fmt"

func writeMetric(name string, value int) error {
	return fmt.Errorf("can't write")
}

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

	if err := writeMetric("num_errors", 0); err != nil {
		fmt.Println(err)
	}
}
