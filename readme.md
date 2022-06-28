
# Build and Run

To add more packages locally, follow these suggestions.  
To not get mad with gitignore, always build giving an ignored extension to binaries (*.o, *.exe), or move them to an ignored folder (/bin)

```go
go mod init myExample
go mod tidy
go run .
```

# Resources
## Tutorials
- https://go.dev/tour
- https://go101.org/
- https://go.dev/doc/tutorial/
- https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go

## Video courses
- [TechWorld with Nana - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=yyUHQIec83I)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https:/www.youtube.com/watch?v=YS4e4q9oBaU)
- [That DevOps Guy - Introduction to Go Programming for beginners](https:/www.youtube.com/watch?v=jpKysZwllVw)

## Exercises 
- https://exercism.org/tracks/go
- https://golangr.com/exercises/
- https://gophercises.com/
- https://www.hackerrank.com/domains/algorithms

## Suggested path (WIP) 
**Use [Go Tour](https://go.dev/tour) as the base lessons order.** While the following list should be used as an extension of the arguments for a deeper understanding:
- Hello world!
    - https://exercism.org/tracks/go/exercises/hello-world
- Modules and packages
    - https://go.dev/doc/tutorial/create-module
    - https://go101.org/article/packages-and-imports.html
- Strings
    - https://www.calhoun.io/6-tips-for-using-strings-in-go/
- Slices
    - [golangbot.com/arrays-and-slices](https://golangbot.com/arrays-and-slices/)
    - [go.dev/blog/slices-intro](https://go.dev/blog/slices-intro)
    - https://pkg.go.dev/builtin#append
    - [growslice() 1.18 code algorithm](https://github.com/golang/go/blob/dcdb19874ff3699e60e41e6b74757b37c4d99b0f/src/runtime/slice.go#L166)
- Maps
- Closures
    - https://go.dev/tour/moretypes/25
    - https://gobyexample.com/closures
    - [Fibonacci exercise](https://go.dev/play/p/MdTuaDKTvzc)
- Interfaces
    - [Go tour - interfaces](https://go.dev/tour/methods/9)
    - [Go101 - Interfaces](https://go101.org/article/interface.html) - Until value boxing chapter
    - https://pkg.go.dev/builtin#any
    - [Print code](https://cs.opensource.google/go/go/+/refs/tags/go1.18.2:src/fmt/print.go;l=273) - fast view
    - [Interface specification](https://go.dev/ref/spec#Interface_types)

- Value boxing, polymorphism, reflection
    - [Go101 Interfaces](https://go101.org/article/interface.html) - From value boxing 
    - https://pkg.go.dev/reflect

- [Memory model](https://go.dev/ref/mem)
    - Profiling 
        - https://pkg.go.dev/runtime/pprof#pkg-overview
        - [HTTP profiler](https://pkg.go.dev/net/http/pprof)
    - [How to dump goroutine stacktraces?](https://stackoverflow.com/a/19094539/3673430)

- Benchmarking
    - https://blog.logrocket.com/benchmarking-golang-improve-function-performance/

- Concurrency and Routines
    - https://go.dev/tour/concurrency/1
    - https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency
    - https://go101.org/article/channel-use-cases.html

## Books
- https://assets.digitalocean.com/books/how-to-code-in-go.pdf

