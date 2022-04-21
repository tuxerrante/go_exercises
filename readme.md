
# Build and Run

To add more packages locally, follow these suggestions.
**Important**: to not get mad with gitignore, always build giving an ignored extension to binaries (*.o, *.exe), or move them to an ignored folder (/bin)
```go
go mod init example/functions
go mod tidy
go mod edit -replace example/functions=../functions


go run example/functions
```

# Resources
## Tutorials
- https://go.dev/tour/methods/20
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

## Books
- https://assets.digitalocean.com/books/how-to-code-in-go.pdf

