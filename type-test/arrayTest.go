package main

import "fmt"

// convert func name in "exported name" with capital letter to call it from another package
func arrayTest() {

	a_1 := [...]int{10: -1}
	fmt.Printf("\n> a_1 [...]int{10: -1} > %v", a_1)

	a_2 := [5]int{10, 20, 30, 40, 50}
	s_2 := append(a_2[:3], a_2[4:]...)
	fmt.Printf("\n> s_2 := append(a_2[:3], a_2[4:]...) > %v", s_2)

}
