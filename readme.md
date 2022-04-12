
To add more packages locally:
```go
go mod init example/functions
go mod tidy
go mod edit -replace example/functions=../functions

go build . && go run example/functions
```
# Resources
- https://go.dev/tour/methods/20
- https://go101.org/
- https://go.dev/doc/tutorial/
- https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go

- https://exercism.org/tracks/go
- https://golangr.com/exercises/
- https://gophercises.com/
