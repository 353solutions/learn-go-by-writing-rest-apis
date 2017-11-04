# Testing

Go comes with a simple test suite. We'll start by writing a simple test and
placing it in a file called `httpd_test.go`.

The `go build` tool ignores files ending with `_test.go` in the normal build process.
While the `go test` tool looks for them and executes every test function in
them.

## Simple Test

Let's start simple, write the following in `httpd_test.go`

```go
package main

import (
	"testing"
)

func TestHome(t *testing.T) {
}
```

To be a test function, it should start with `Test` and except one parameter of
type `*testing.T`.

Once you have that, execute

    go test -v

You should see your first passing test!

## Handler Test

To test our HTTP server we have two options. We can either run it as a process
and then issue HTTP calls to it, or we can test the handlers directly.

We'll go with the second approach and use the `net/http/httptest` package. Which
lets use create new requests and `http.ResponseWriter` that records what was
written to it.

If you look in the [documentation of httptest][httptest] you'll see:

```go
func NewRequest(method, target string, body io.Reader) *http.Request
```

The `net/http/httptest` defines `NewRequest`. We don't have constructors like in
other OO packages. In Go we define `New<TYPE>` function that returns a new
object.

[httptest]: https://golang.org/pkg/net/http/httptest/

### Variable Definitions

Before we go on, we need to see how we can define new variables. We can either
use `var` as in

```go
var i int
```

`i` is of type `int` (which is 64 bits on my machine), it will be initialize to
the zero value of its type, which is the number 0.

We can also assign a value when we define the variable

```go
var i int = 42
```

The last form, and the most common one, it to let the Go compiler infer the
variable type from the value.

```go
i := 42
```

We can also assign multiple variables in the same line

```go
i, s := 42, "hello"
```

Functions in Go can return multiple values, and then we can use this multiple
assignment.

```go
val1, val2 = someFunction()
```

It's very common in Go that the last value return is an error code. We don't use
exceptions as in Python/C++/Java...

### `if` Statement

Go has the usual `if` statement.

```go
if i > 0 {
    fmt.Println("i is big")
}
```

We don't need to place `()` around the condition.

We can combine conditions with `||` for `or`, `&&` for `and` and negate them with
`!`

```go
if i > 0 && i < 10 {
    fmt.Println("i is just right")
}
```

Go has the usual comparison operators: `==`, `!=`, `>`, `>=`, `<` and `<=`.

We can also assign and check in one `if` statement. This is very common combined
with error checking.

```
if err := writeMetric("num_errors", 0); err != nil {
	fmt.Println(err)
}
```

## Exercise

Change `TestHome` to call `handler` and check that we get `שלום Gophers` back.

Hint: See `ResponseRecorder` [example][resp-rec].

[resp-rec]: https://golang.org/pkg/net/http/httptest/#example_ResponseRecorder

[Solution](httpd_test.go)
