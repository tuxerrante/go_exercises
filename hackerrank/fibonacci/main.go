package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */
const mod = 1000000000

// Keep track of previous numbers
//
//	No direct infers from sizes since there is the skip parameter
var currentNumber int = 1
var lastResultTime time.Time = time.Now()

func ModuloFibonacciSequence(requestChan chan bool, resultChan chan int) {
	for <-requestChan {
		result := moduloFibonacci(currentNumber)

		currentNumber += 1

		// Send responses no earlier than 10 ms after the previous one
		// Duration represents the elapsed time between two instants as an int64 nanosecond count
		for time.Since(lastResultTime) < time.Millisecond*10 {
			continue
		}
		lastResultTime = time.Now()
		resultChan <- result
	}
}

// Depending on how you want to start n1,n2 could be 0,1 or 1,0
func moduloFibonacci(target int) int {
	n1, n2 := 1, 0
	for i := 0; i < target; i++ {
		n1, n2 = n2, n1+n2
	}
	return (n1 + n2) % mod
}

// Not efficient
func moduloFibonacciRecursive(currentNumber int) int {
	if currentNumber == 0 {
		return 0
	} else if currentNumber == 1 {
		return 1
	}

	return (moduloFibonacciRecursive(currentNumber-1) + moduloFibonacciRecursive(currentNumber-2)) % mod
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	file, err := os.Open("./tests/test_02.in")
	if err != nil {
		panic("File access error!")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	skip := int32(skipTemp)
	log.Printf("skip = %d\n", skip)

	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	total := int32(totalTemp)
	log.Printf("total = %d\n", total)

	resultChan := make(chan int)
	requestChan := make(chan bool)

	go ModuloFibonacciSequence(requestChan, resultChan)

	for i := int32(0); i < skip+total; i++ {
		start := time.Now().UnixNano()

		log.Printf("Sending to request channel..\n")
		requestChan <- true

		new := <-resultChan
		if i < skip {
			continue
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff < 3 {
			fmt.Println("Rate is too high")
			break
		}
		fmt.Println(new)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
