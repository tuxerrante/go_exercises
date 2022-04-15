package main
import (
	"fmt"
	"strings"
	"math"
)

func main (){
	var edge int = 3
	fmt.Printf("> The cube of %d is %d\n", edge, cube(edge))

	// s := noReturn("foo") => noReturn("foo") (no value) used as value
	
	fmt.Printf("> Sum = %d\n", sumNumbers())
	fmt.Printf("> Sum = %d\n", sumNumbers(1, 2, 10))


	// Anonymous function
	x := func(s string) string {
		return strings.ToUpper(s)
	}
	fmt.Printf("%T \n", x)
	fmt.Println(x("Joe"))

	// --- High Order function
	fmt.Println()
	radius := 3.4
	// Simulate a menu choice like
	// fmt.Scanf("1. Area\n 2. Perimeter")
	inputChoice := 2
	printCircleResult(radius, selectCircleChoice(inputChoice))

	//
	fmt.Println()
	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Printf("> Swap(1, 2)= %v,%v\n", a, b)
	fmt.Println(swap(a,b) )
}

func cube (edge int) int {
	return edge*edge*edge
}

// if no return value is returned, it is not even nil
// x := noReturn("foo") will fail!
func noReturn(s string) {
	fmt.Println(s)
}

// variadic function
// ... elypsis for creating an argument which is an int slice
func sumNumbers(numbers ...int) (sum int) {
	for _, num := range numbers{
		sum += num
	}
	return
}

// https://go.dev/blog/defer-panic-and-recover
// Deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:
func c() (i int) {
    defer func() { i++ }()
    return 1
}


// --- High order function
// 		calcFunction will be the result of the map query created by selectCircleChoice
func printCircleResult(radius float64, calcFunction func(float64) float64){
	result := calcFunction(radius)
	fmt.Println("> Circle Result:", result)
}
func selectCircleChoice (query int) func(float64) float64 {
    choice := map[int]func(r float64) float64{
        1: calcArea,
        2: calcPerimeter,
    }
    return choice[query]
}
func calcArea(r float64) float64{
    return math.Pi*r*r
}
func calcPerimeter(r float64) float64{
    return math.Pi*r
}

// --- Naked return should be used only in short functions
func swap(x, y int) (a, b int) {
	a = y
	b = x
	return
}