package main

import "fmt"

func main() {
	s := [3]int{10, 20, 30}
	fmt.Println(s)

	modify(s[:])
	fmt.Println("Will change the reference, since it is passed as slice", s)

	modifyArray(s)
	fmt.Println("Won't Change the reference, since it is passed as an array", s)

	ascii := map[string]int{}
	ascii["A"] = 65
	modifyMap(ascii)
	fmt.Println("Will change the reference, since it is passed as a map", ascii)

}

func modify(slice []int) {
	for i := range slice {
		slice[i] -= 5
	}
}

func modifyMap(m map[string]int) {
	m["A"] = 100
}

func modifyArray(arr [3]int) {
	for i := range arr {
		arr[i] -= 6
	}
}
