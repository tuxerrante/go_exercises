package runesExamples

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)
func PrintRunesExamples(){

	var counter map[rune]int = make(map[rune]int)	// map to count runes
	var invalid int 		 = 0
	var utf8Lengths [utf8.UTFMax +1]int 			// array of runes lengths: encoding length hs range only from 1 to utf8.UTFMax (which has the value 4)
	
	const unicode_input string = "1234 5678 ° 27 B 15 F 14 é 13 A 10 < 5 & 5 D 4 ( 4 + 3 \""

	/*
		Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text. 
		Successive calls to the Scan method will step through the 'tokens' of a file, skipping the bytes between the tokens.
		in := bufio.NewScanner(os.Stdin)
	*/

	/*
		Reader is the interface that wraps the basic Read method.
		When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read.
	*/
	in := bufio.NewReader(strings.NewReader(unicode_input))
	
	for {
		currentRune, runeSize, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR> counter:%v > %v\n", counter, err)
			os.Exit(1)
		}
		if currentRune == unicode.ReplacementChar && runeSize == 1 {
			invalid++
			continue
		}
		counter[currentRune]++
		utf8Lengths[runeSize]++
	}

	fmt.Printf("Rune\tCount\n")
	for r,c := range counter {
		fmt.Printf("%v\t%d\t%#U\t%s\n", r, c, r, strconv.QuoteRune(r))
	}

	fmt.Printf("\nLen\tCount\n")
	for r,size := range utf8Lengths {
		if size > 0 {
			fmt.Printf("%v\t%d\n", r, size)
		}
	}

	if invalid > 0 {
		fmt.Printf("Found %d invalid Unicode code points", invalid)
	}
}