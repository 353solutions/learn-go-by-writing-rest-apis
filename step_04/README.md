# Setting & Getting Values

## Maps

maps (also known as hash tables) are a data structure that maps keys to values.

We can create maps in two ways, either with `make`

```go
m := make(map[string]int)
m["x"] = 1
m["y"] = 2
m["z"] = 3
```

or verbatim

```go
m := map[string]int{
	"x": 1,
	"y": 2,
	"z": 3,
}
```

We set/change and get values with `[]`

```go
m["x"] = 1
m["x"] = 2
val := m["x"]
```

And we can check if a keys exists with

```go
_, ok := m["x"]
if ok {
	fmt.Println("x in m")
}
```

We delete values with `delete`

```go
delete(m, "x")
```

You can see more about maps [here][maps.go]

## Slicing

To get parts of string (and other data structures) we can use slicing. The
syntax of slicing if `[start:stop]` where `start` and `stop` are indices. Slices
are "half-open" meaning we'll get the first index but not the last. We can omit
the `start` or `end` and Go will fill then in for us.

```go
book := "the colour of magic"

fmt.Println(book[4:10])  // "colour"
fmt.Println(book[11:]) // "of magic"
fmt.Println(book[:3]) // "the"
```

## Byte Slices

When we read from files (or sockets, or ...) we get a byte slice. It's an array
of bytes and the type is `[]byte`. We'll talk about slices and arrays later.

## defer

Go have an automatic memory management (known as garbage collector or GC for
short). However in our program we sometimes user other resources and we'd like
to make sure they are closed when we're done with them.

In C++/Java have `finally`, in Python we have `with` and in Go we have `defer`.

`defer` get a function call and will execute is ones the current function
exists.

```go
func cleanup() {
	fmt.Println("cleanup")
}

func caller() {
	defer cleanup()

	fmt.Println("caller code")
}
```

When we run this we'll see

```
caller code
cleanup
```

`defer`ed are called in reverse order of invocation. Here's an example of
making sure we closed a file once we opened it.

```go
file, err := os.Open("defer.go")
if err != nil {
	fmt.Printf("error: can't open - %s\n", err)
	return
}
defer file.Close()

// Work with file here
```


## Exercise

Add an option to set and get values from our server.

* To set a value make a POST call `/db/<key>` with value as data
* To get a value make a GET call to `/db/key`

Hint: To route everything under `/db` to the handler, mount it on `/db/` (e.g.
`http.HandleFunc("/db/", dbHandler)`)


### Testing

Set Value

    curl -XPOST -d1 http://localhost:8080/db/x

Get Value

    curl http://localhost:8080/db/x

[Solution](httpd.go)
