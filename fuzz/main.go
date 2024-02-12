package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := ReverseString(input)
	doubleRev, _ := ReverseString(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func ReverseString(s string) (string, error) {
	r := []rune(s)

	log.Printf("Reverting string '%s'", s)

	if !utf8.ValidString(s) {
		return "", fmt.Errorf("cannot reverse invalid UTF-8 string %q", s)
	}

	for left, right := 0, len(r)-1; left < len(r)/2; left, right = left+1, right-1 {
		r[right], r[left] = r[left], r[right]
	}

	return string(r), nil
}
