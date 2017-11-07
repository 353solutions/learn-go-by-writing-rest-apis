# gorilla/mux


## go get

Go is a simple language by design, and the design of `net/http` reflects that.
Sometimes we'd like easier support for routing and other tasks. For this we can
install a third-party package.

We're going to use [gorilla/mux][mux] to simplify out routing.

To install third-party package we'll be using `go get`. Run the following
command in a terminal/command prompt:

```
go get github.com/gorilla/mux
```

This will fetch (using `git`) and install the `gorilla/mux` package under your
current `$GOPATH`. You can start importing the package in your code now.

If the Go code finds out that a package you install depends on other packages -
it will install these dependencies as well. There are several external tools
for managing dependencies and [dep][dep] is emerging as the official one - but
it's still in beta.

[dep]: https://github.com/golang/dep
[mux]: http://www.gorillatoolkit.org/pkg/mux

## mux

A mux (multiplexer) is a router that routes HTTP paths to handler. You can use
regular expression in paths and also give names to parts of the path. Let's see
an example.

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	fmt.Fprintf(w, "Hello %s\n", name)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/hi/{name}", hiHandler)
	http.ListenAndServe(":8080", rtr)
}
```

If we run this

    go run mux.go

And the hit the API 

    curl http://localhost:8080/hi/There

we'll get back "Hello There"

`gorilla/mux` also allows you to define the HTTP methods/verbs (GET, POST ...)
that are allowed to each handler.

## Exercise

Use `gorilla/mux` in our sever to figure out the key name and split `dbHandler`
to a `getHandler` (GET request) and `setHandler` (POST request).


[Solution](httpd.go)
