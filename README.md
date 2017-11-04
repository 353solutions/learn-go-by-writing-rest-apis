# Learn Go by Writing REST APIs

Miki Tebeka <miki@353solutions.com>

## Abstract
In this hands-on course, learn get started with the [Go programming
language][go]. Learning will be done by writing some REST APIs. This class is
for people who would like to get started with Go.

[go]: https://golang.org

## Chapters

* In [Setup][setup] we'll make sure you have all the prerequisites
* We'll write the customary ["Hello World" program][hello-world]
    - Program structure, import path, types
* Now we can add some [tests][tests]
    - Testing, declarations, `if` statement
* Next we'll add [ping & time][ping] handlers
    - HTTP protocol, REST overview
* After that we'll [set and get keys][setget]
    - map, slicing, byte arrays, defer
* Then we'll return all the [keys][keys]
    - slices, `for` loop, empty interface, JSON
* We'll use [gorilla/mux][mux] to simplify routing
    - `go get`
* Finally we'll support [concurrency][sync]
    - goroutines, locks, struct

[setup]: step_00/README.md
[hello-world]: step_01/README.md
[tests]: step_02/README.md
[ping]: step_03/README.md
[setget]: step_04/README.md
[keys]: step_05/README.md
[mux]: step_06/README.md
[sync]: step_07/README.md

## Go Wild

The best way to learn a language is to find a problem and solve it using a
language. If the problem is something that "itches" you - even better.

Here are some idea on how to improve our we server:

* Add command line argument for the port to listen on
* Support multiple databases
* Instead of a `sync.Mutex` to guard our data structure. Use a goroutine that
  will listen on a channel for commands and return result over a channel
* Implement more commands, see [redis commands][redis] for inspiration
* Test everything
* ...

[redis]: https://redis.io/commands

## Examples

List of examples used in the lessons:

* [cond.go](examples/cond.go) - Conditionals
* [defer.go](examples/defer.go) - defer
* [empty-iface.go](examples/empty-iface.go) - The empty interface
* [for.go](examples/for.go) - For loops
* [goroutines.go](examples/goroutines.go) - goroutines
* [json.go](examples/json.go) - JSON encoding
* [maps.go](examples/maps.go) - maps
* [mux.go](examples/mux.go) - gorilla/mux
* [slices.go](examples/slices.go) - Slices
* [struct.go](examples/struct.go) - Structs
* [vars.go](examples/vars.go) - Variable declaration

## License

[CC BY-NC-SA 4.0][license]

[license]: https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode
