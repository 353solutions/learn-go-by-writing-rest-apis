# HTTP & Routing

## HTTP Overview

HTTP is a textual protocol over TCP/IP. An HTTP call looks like:

```
POST /x?ttl=10 HTTP/1.1
Host: httpbin.org
Connection: close
Content-Length: 5

hello
```

The first line (request line) states the request:
* `POST` is the verb or method
* `/x?ttl=10` is the path
* `HTTP/1.1` is the protocol

After that we have a bunch of headers in `key: value` format. Then an empty
line and (optional) body.

### URLs

Let's assume we call `https://httpbin.org/get?x=1?y=2`, then

* `https` is the protocol
* `httpbin.org` is the host
* `/get` is the path
* `x=1?y=2` are the parameters (or query)

## REST Overview

Most operations on data can be summed with `CRUD` - Create, Retrieve, Update and
Delete.

In REST we make HTTP calls to a server to and use HTTP verbs to describe which
operation we'd like to do. The mapping of CRUD to REST is

* **C**create -> POST
* **R**etrieve -> GET
* **U**pdate -> PUT/PATCH
* **D**elete -> DELETE

We can find out the HTTP method from the request object `Method` field.

When we make an HTTP call with `curl` we can use the `-X` switch to define the
method.

     curl -XPOST https://httpbin.org/post

We can pass data in the request body. The request data is available in the
`Body` field of the request and it's an `io.Reader` (since request body can be
very large). 

When we make an HTTP call with curl we can either pass the data using the `-d`
switch or pass a file if the data starts with `@`.

     curl -XPOST -d hello https://httpbin.org/post
     curl -XPOST -d@httpd.go https://httpbin.org/post

The resource we're accessing is defined by the path after the host name. For
example in `http://localhost:8080/a/b` the path will be `/a/b`.

In our handler we can get the path from `r.URL.Path`

We can also pass parameters in HTTP calls. We pass them after `?` in the URL and
by name separated by `&`, for example

    curl "https://httpbin.org/get?x=1&y=2"  

In our handler we can access these parameters via `r.URL.Query`

### HTTP Errors

When we return an HTTP response, we return an HTTP status code with the
response. There are [many status codes][http-status] codes available. Basically
200s means OK, 300s means redirect, 400s means the request was bad and 500s mean
there was some error in the server.

In our handlers, we can use `w.WriteHeader` to set the status. There also
[http.Error][http-error] to set the status and return an answer.

[http-status]: https://www.flickr.com/photos/girliemac/sets/72157628409467125/
[http-error]: (https://golang.org/pkg/net/http/#Error)


## Exercise

Add two HTTP end points to our server.

* `/ping` should accept GET request and return `PONG`
    * If ping gets a `msg` parameter it should return it instead of `PONG`
* `/time` should accept the a GET request and return the current time in RFC3339
  format

Hint: The [time][time] package

[Solution](httpd.go)

[time]: https://golang.org/pkg/time/
