package go-types
import "fmt"

/*
	C types is the memory structures of C values are transparent. 
	Each C value in memory occupies one memory block (one continuous memory segment). 
	However, a value of some kinds of Go types may often be hosted on more than one memory block.

	https://go101.org/article/value-part.html#interface-structure
	
*/