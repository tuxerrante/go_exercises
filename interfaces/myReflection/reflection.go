package myReflection

import (
	"fmt"
	"log"
)

/*
	built-in reflections are achieved with type assertions and type-switch control flow code blocks.

	Type assertion.
	There are four kinds of interface-value-involving value conversion cases in Go:
	COMPILE TIME CHECK:
	- convert a non-interface value to an interface value,
		where the type of the non-interface value must implement the type of the interface value.
	- convert an interface value to an interface value,
		where the type of the source interface value must implement the type of the destination interface value.

	RUNTIME CHECK:
	- convert an interface value to a non-interface value,
		where the type of the non-interface value must implement the type of the interface value.
	- convert an interface value to an interface value,
		where the type of the source interface value doesn't implement the destination interface type,
		but the dynamic type of the source interface value might implement the destination interface type.

	- https://go101.org/article/interface.html#reflection
	- https://go101.org/article/reflection.html
	- https://golang.org/pkg/reflect/
*/

/*
	In the case of T being a non-interface type, if the dynamic type of i exists and is identical to T, then the assertion will succeed,
*/
func NonInterfaceTypes(){
	
	var i interface{} = 123

	// asserted type is a non-interface int value
	n, ok := i.(int)
	if ! ok {
		log.Println("nonInterfaceTypes(): can't complete assertion from int to int!")
	}
	describe(n)
		
	n = i.(int)		// we can now assign the value
	describe(n)

	// let's try with a float
	n2, ok := i.(float64)
	if ! ok {
		log.Println("nonInterfaceTypes(): can't complete assertion from int to float64!")
	}
	describe(n2)
	// will panic: n2 = i.(float64)

}


type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}
func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

type fakeString struct {
	a string
}
func (fakeString) Write(buf []byte) (int, error){
	return 1, nil
}

func InterfaceTypes() {
	
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"

	// Can I assert from a string to a type interface with a dynamic of string
	var ss interface{} = fakeString{}
	ss = fakeString{ a: "abc"}
	fmt.Println(ss)
	ss2, ok2 := ss.(string)
	if ! ok2 {
		fmt.Println("> KO")
	}
	describe(ss2)

	// Now the dynamic type of y is "string".
	var w Writer
	var ok bool

	// Type DummyWriter implements both
	// Writer and interface{}.
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// The dynamic type of y is "string",
	// which doesn't implement Writer.
	w, ok = y.(Writer)
	fmt.Println(w, ok) //  false
	// will panic:  
	// w = y.(Writer)

}


func describe(i interface{}) {
	fmt.Printf("%T\t%+v\n", i,i)
}