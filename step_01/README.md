# Hello World

Let's start by writing the customary "Hello World" program.

[hw.go](hw.go)


```go
package main

import "fmt"

func main() {
	fmt.Println("שלום Gophers")
}
```

## Running

Go compile our program run

    go build -o hw

The `-o hw` tells the `go` tool the name of the executable to create. If you
don't specify it it'll be renamed in the name of the current directory.

(On Windows the executable will be called `hw.exe`)

Now we can run it

    ./hw


## Line by Line

Let's go over the code line by line.


### `import "fmt"`

Go is
