package main

/*
	https://go101.org/article/interface.html
	https://go.dev/tour/methods/14
	https://go.dev/ref/spec#Interface_types
	https://go-book.readthedocs.io/en/latest/interfaces.html
	-----------------------------------------------------------------------------
	Implementations are all implicit in Go.
	The compiler does not require implementation relations to be specified in code explicitly.
*/
import (
	"fmt"
	"myInterfaces/myReflection"
)

func main() {
	/*
		// Every type that is a member of the type set of an interface implements that interface.
		// Any given type may implement several distinct interfaces.
		var i interface{}
		describe(i)


		// the predeclared type any is an alias for the empty interface.
		// https://pkg.go.dev/builtin#any
		var i2 any
		describe(i2)

		var uns_int uint8 = 200
		describeUnsigned(uns_int)

		// use non basic interface to restrict types, not for conversions
		var uint1 uint = 10
		var uint2 uint = 32
		fmt.Println("> positiveSum(uint1, uint2)", positiveSum(uint1, uint2) )

		// Similar test with an "underlying" type []byte
		type Bytes []byte
		var myBytes Bytes = make(Bytes, 3)
		myBytes[0] = 42
		fmt.Println("> firstByte(myBytes)", firstByte(myBytes) )

		type Bytes2 Bytes
		var myBytes2 Bytes2 = make(Bytes2, 3)
		myBytes2[0] = 43
		fmt.Println("> firstByte(myBytes)", firstByte(Bytes(myBytes2)) )
	*/

	// REFLECTION
	fmt.Println("---\tREFLECTION")
	myReflection.NonInterfaceTypes()
	myReflection.InterfaceTypes()

}

// Describe
func describe(i interface{}) { // print interface{} value and type
	fmt.Printf("(%v, %T)\n", i, i)
}

// Interfaces whose type sets can be defined entirely by a list of methods are called basic interfaces.
type reallyBasic interface {
	isBasic() bool
}

type basicIO interface {
	Write() (int, error) // |__ Type sets
	Read() (int, error)  // |
	reallyBasic
}

/*
Interfaces that are not basic may only be used as type constraints,
or as elements of other interfaces used as constraints.
*/
type UnsignedInt interface { // non-basic: has a type element (union)
	uint | uint8 | uint16 | uint32 | uint64
}

type UnsignedStuff interface {
	UnsignedInt // embedding interface UnsignedInt in UnsignedStuff
	IsNice() bool
}

func positiveSum[T UnsignedInt](a, b T) T {
	return a + b
}

/*											error -> interface contains type constraints
func describeUnsigned(ui UnsignedInt) {
	fmt.Println(ui)
}
*/

/*
General interfaces
An interface representing all types with underlying type []byte
In a term of the form ~T, the underlying type of T must be itself, and T cannot be an interface.
*/
type AnyByteSlice interface { // non-basic: This interface embeds an approximation type.
	~[]byte
	// ~error   					illegal: error is an interface
}

func firstByte[T AnyByteSlice](t T) T {
	if t != nil {
		return t[:1]
	}
	return nil
}
