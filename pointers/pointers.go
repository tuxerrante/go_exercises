package main
import (
	"fmt"
)
func main (){

	i := 10
	fmt.Printf(" Var address type %T and value %v\n", &i, &i)
	fmt.Printf(" Var dereference type %T and value %v\n", *(&i), *(&i))

	var pointer_i *int = &i
	fmt.Printf(" Pointer type %T and value %v\n", pointer_i, pointer_i)

}