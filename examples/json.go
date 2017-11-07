// JSON encoding
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	data := map[string]interface{}{
		"name": "daffy",
		"age":  80,
	}

	enc := json.NewEncoder(os.Stdout) // Encode to standard output
	if err := enc.Encode(data); err != nil {
		fmt.Printf("error: can't encode - %s\n", err)
	}
}
