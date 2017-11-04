# Concurrency

Go has excellent support for concurrency. It uses two primitives to do that -
goroutines and channels.

## goroutines

goroutines are "light weight" threads. They start with a 2Kb stack that grows
with need. This allows us to run thousands of goroutines on a single machine
without any issue. Every time our we server get a new connection, it spins a new
goroutine to handle it. 

Here's a simple [example](goroutines.go)

```go
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
```

Go's basic data structures are not "goroutine safe", which means that when you
modify them from two goroutines at the same time - something bad will happen.

To guard our data structures we can use [sync.Mutex][mutex] which ensures that
only one goroutines modifies a data structure at a point in time.

There are other ways to do concurrency in Go, such as atomic changes and
channels. But we won't get there this time.

[mutex]: https://golang.org/pkg/sync/#Mutex


## Structs

Sometimes we'd like to define out own data type. In this case we'll use a
`struct` which is a set of name fields, each with it's own type.

Let's define a [Loon][loon] struct that will hold name and age.

```go
type Loon struct {
	Name string
	Age  int
}
```

We can create structs by passing fields in orders.

```go
	elmer := Loon{"Elmer", 80}
```

or we can specify the field names

```go
	daffy := Loon{
		Age:  80,
		Name: "Daffy",
	}
```

We don't have constructors in Go, we simply write a function that creates a new
struct and return it. This function is  usually called `New<Type>`

```go
func NewLoon(name string, age int) *Loon {
	return &Loon{name, age}
}
```

We return a pointer to the object we created, so we can reference it and not
copy it over.

People coming from C/C++ will see a but in `NewLoon` since it returns a pointer
to something allocated on the stack. The Go compiler does what's called "escape
analysis" and sees that we return a pointer to the new allocated `Loon` so it'll
create it on the heap.

`struct`s can also have methods, embed others and much more. But we won't get
into that.

[loon]: https://en.wikipedia.org/wiki/Looney_Tunes

## Exercise

Change our `db` to be a struct that holds the map and also a `sync.Mutex`.
Every time we access the data, `Lock` the mutex, make sure to `Unlock` it with
`defer`

[Solution](httpd.go)
