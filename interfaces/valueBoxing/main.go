package main
/*
	https://go101.org/article/interface.html
*/
import "fmt"

type Aboutable interface {
	About() string
}

// Type *Book implements Aboutable.
type Book struct {
	name string
}
func (book *Book) About() string {
	return "Book: " + book.name
}

func main() {
	// A *Book value is boxed into an
	// interface value of type Aboutable.
	var a Aboutable = &Book{"Go 101"}
	fmt.Println(a) // &{Go 101}

	// i is a blank interface value.
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}

	// Aboutable implements interface{}.
	i = a
	fmt.Println(i) // &{Go 101}
}