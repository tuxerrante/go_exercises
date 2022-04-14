```go
go mod init example/functions
go mod tidy
go mod edit -replace example/functions=../functions

go build -o main.o . && go run example/functions

```
This also works
```go
go build -o main.o example/functions && go run example/functions
```