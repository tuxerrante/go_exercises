package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	files, err := os.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
