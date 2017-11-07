# List of Keys

## Slices

A Go [slice][slice] is a sequence of typed data

```go
names := []string{"bugs", "daffy", "elmer"}
```

We can use slicing on slices

```go
fmt.Println(names[1:])
```

To add values, use [append][append]

```go
names = append(names)
```

[append]: https://golang.org/pkg/builtin/#append
[slice]: https://blog.golang.org/go-slices-usage-and-internals

## `for` Loops

The only iteration mechanism in Go is the `for` loop.

We can use `for` like we do in many other languages

```go
for i := 0; i < 3; i++ {
	fmt.Println(i)
}
```

We can use it as a `while` loop

```go
n := 3
for n > 0 {
	fmt.Println(n)
	n--
}
```

Also as an infinite loop

```go
j := 0
for {
	fmt.Println(j)
	j++
	if j > 2 {
		break // Stop the loop
	}
}
```

To help avoid making `off-by-one` errors, we have `range`.

`range` in single value context will go over indices

```go
s := "abc"
for i := range s {
	fmt.Println(i)
}
```

`range` in double value context will go over indices & values

```go
for i, c := range s {
	fmt.Println(i, c)
}
```

## The Empty Interface

Go is a typed language. In some ways it's more strict than C/Java - for example
you can't divide a integer with a float. However sometimes we'd like a variable
that holds any type (`object` in Java, `void *` in C/C++ ...). In Go this type
is the empty interface, which is written as `interface{}`.

```go
var i interface{}

i = 10

// Print value and type
fmt.Printf("%v (%T)\n", i, i) // 10 (int)

i = "hello"
fmt.Printf("%v (%T)\n", i, i) // hello (string)
```

## JSON

[JSON][json] is a very common serialization format used in REST APIs. Go has
JSON encoders and decoders that will take a Go data structure and convert it
to/from a byte slice.

Since JSON objects (Go maps or structs) can hold different types, we'll be using
`interface{}`

```go
data := map[string]interface{}{
	"name": "daffy",
	"age":  80,
}

enc := json.NewEncoder(os.Stdout) // Encode to standard output
if err := enc.Encode(data); err != nil {
	fmt.Printf("error: can't encode - %s\n", err)
}
```

[json]: http://www.json.org/


## Exercise

Add a `/keys` entry point which will return all the keys in the database as a
JSON array.


### Testing

    curl http://localhost:8080/keys

[Solution](httpd.go)
