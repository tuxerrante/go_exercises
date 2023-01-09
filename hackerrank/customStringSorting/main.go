package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
An odd length string should precede an even length string.
If both strings have odd lengths, the shorter of the two should precede.
If both strings have even lengths, the longer of the two should precede.
If the two strings have equal lengths, they should be in alphabetical order.

The function is expected to return a STRING_ARRAY.
The function accepts STRING_ARRAY strArr as parameter.
*/
func customSorting(strArr []string) []string {

	var evenStringLen map[string]int = map[string]int{}
	var oddStringLen map[string]int = map[string]int{}
	var evenStrings []string = []string{}
	var oddStrings []string = []string{}

	for _, s := range strArr {
		// Separate strings by length in two different arrays
		log.Printf("Reading %q\n", s)
		lenStr := len(s)
		isOdd := lenStr%2 != 0

		// ODD STRINGS
		if isOdd {
			// Save the length of that string in a map to avoid compute it again for all strings
			oddStringLen[s] = lenStr
			log.Printf("%q is of size %d and it is ODD\n", s, lenStr)

			if len(oddStrings) == 0 {
				oddStrings = append(oddStrings, s)
				log.Println(oddStrings, oddStringLen)
				continue
			}

			// compare odd strings length and alphabetically
			for index, currentOddString := range oddStrings {
				if index == len(oddStrings)-1 {
					oddStrings = append(oddStrings, s)
				} else if oddStringLen[s] < oddStringLen[currentOddString] {
					oddStrings = append(oddStrings[:index+1], oddStrings[index:]...)
					oddStrings[index] = s
					break
				} else if oddStringLen[s] == oddStringLen[currentOddString] {
					// strings have same length
					if strings.Compare(s, currentOddString) == 1 {
						// s is greater, I want currentOddString first
						oddStrings = append(oddStrings[:index+1], oddStrings[index:]...)
						oddStrings[index+1] = s
					} else {
						// s is smaller, I want it before currentOddString
						oddStrings = append(oddStrings[:index+1], oddStrings[index:]...)
						oddStrings[index] = s
					}
					break
				}
				log.Println(oddStrings, oddStringLen)
			}
			// EVEN STRINGS
		} else {
			evenStringLen[s] = lenStr
			log.Printf("%q is of size %d and it is EVEN\n", s, lenStr)

			if len(evenStrings) == 0 {
				evenStrings = append(evenStrings, s)
				log.Println(evenStrings, evenStringLen)
				continue
			}

			// compare even strings length and alphabetically
			for index, currentEvenString := range evenStrings {
				if index == len(evenStrings)-1 {
					evenStrings = append(evenStrings, s)
				} else if evenStringLen[s] < evenStringLen[currentEvenString] {
					evenStrings = append(evenStrings[:index+1], evenStrings[index:]...)
					evenStrings[index+1] = s
					break
				} else if evenStringLen[s] == evenStringLen[currentEvenString] {
					// strings have same length
					if strings.Compare(s, currentEvenString) == 1 {
						// s is greater, I want currentEvenString first
						evenStrings = append(evenStrings[:index+1], evenStrings[index:]...)
						evenStrings[index+1] = s
					} else {
						// s is smaller, I want it before currentEvenString
						evenStrings = append(evenStrings[:index+1], evenStrings[index:]...)
						evenStrings[index] = s
					}
					break
				}
				log.Println(evenStrings, evenStringLen)
			}
		}
	}
	// Append the two results
	return append(oddStrings, evenStrings...)
}

// local test: export OUTPUT_PATH="test/test.out"
func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	// ----- Local testing
	log.SetFlags(log.Ltime)
	file, err := os.Open("test/test_01.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// -----

	reader := bufio.NewReader(file)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var strArr []string

	for i := 0; i < int(strArrCount); i++ {
		strArrItem := readLine(reader)
		strArr = append(strArr, strArrItem)
	}

	result := customSorting(strArr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
