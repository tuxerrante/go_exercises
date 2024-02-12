package main

import (
	"fmt"
	"log"
	"myExample/stupidPackage"
)

func init() {
	log.SetFlags(3)
}

func main() {

	// olly := "Olly"
	// pointingToOlly := &olly
	var voidVar string
	pointingToOlly := &voidVar

	describe(voidVar)
	// describe(pointingToOlly)

	greeting, err := stupidPackage.Hello(pointingToOlly)
	if err != nil {
		log.Print("this is a log!!!!")
		panic(err)
	}
	fmt.Println(greeting)
}

func describe(i any) {
	fmt.Printf("'%+v' %+T\n", i, i)
}
