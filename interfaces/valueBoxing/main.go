package main

/*
	https://go101.org/article/interface.html

	We can view each interface value as a box to encapsulate a non-interface value.
	To box/encapsulate a non-interface value into an interface value, the type of the non-interface value must implement the type of the interface value.
	The type information of the boxed value is also stored in the result (or destination) interface value.
	When a value is boxed in an interface value, the value is called the dynamic value of the interface value.
	An interface value boxing a nil non-interface value still boxes something, so it is not a nil interface value

	The implementation information for each non-interface type and interface type pair will only be built once
	and cached in a global map for execution efficiency consideration.
	The number of entries of the global map never decreases.
	In fact, a non-nil interface value just uses an internal pointer field which references a cached implementation information entry.
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
	// A *Book value is boxed into an interface value of type Aboutable.
	var a Aboutable = &Book{"Go 101"}
	describe(a) // &{Go 101}

	// i is a blank interface value.
	var i interface{} = &Book{"Rust 101"}
	describe(i) // &{Rust 101}

	// Aboutable implements interface{}.
	i = a
	describe(i) // &{Go 101}

	i = "abc"
	describe(i) // abc

	// Clear the boxed value in interface value i.
	i = nil
	describe(i) // <nil>
}

func describe(i interface{}) {
	fmt.Printf("%T\t%+v\n", i, i)
}
