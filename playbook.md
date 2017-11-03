* REST overview
    - CRUD -> REST (C: POST, R: GET, U: PUT/PATCH, D: DELETE)
    - HTTP/JSON (but not only JSON)
    - JSON can be also just "1" (base64 for bytes)
* Project overview - redis like
    - `docker run redis:alpine`
    - `docker exec -it $(docker ps -q) redis-cli`
    - `SET X 17`
    - `GET X`
* Set working environment - GOPATH
    - vendor
    - dep
* Start with PING, TIME
    - net/http
    - encoding/json
* Test
* Start with SET, GET, DELETE, KEYS
    - maps
    - types and `interface{}`
    - `defer resp.Body.Close()`
    - interfaces
    - errors
* APPEND, LLEN, LRANGE
    - slices
    - for loop
    - advanced: negative indices
* INCR, DECR
    - One goroutine that gets command and return channel
* /{db}/{command}
    - gorilla/mux
    - go get
* Command line options with `flag`
    - `-port`
* TTL (advanced)
