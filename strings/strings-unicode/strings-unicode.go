/* 
	https://go.dev/blog/strings
	only string literals are UTF-8. String values can contain arbitrary bytes; 
	string literals always contain UTF-8 text as long as they have no byte-level escapes.

	Go source code is always UTF-8.
	A string holds arbitrary bytes.
	A string literal, absent byte-level escapes, always holds valid UTF-8 sequences.
	Those sequences represent Unicode code points, called runes.
	No guarantee is made in Go that characters in strings are normalized.
*/
package main
import (
	"fmt"
)
func main() {
    const placeOfInterest = `⌘`

    fmt.Printf("plain string: ")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes: ")
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n\n")

	// The Unicode standard uses the term “code point” to refer to the item represented by a single value.
	var codePoint rune = 0x2318
	fmt.Printf("%s\n", codePoint)
	fmt.Printf("%+x\n", codePoint)
	fmt.Printf("%q\n", codePoint)
	fmt.Printf("%+q\n", codePoint)

	/* A for range loop, decodes one UTF-8-encoded rune on each iteration. 
		Each time around the loop, the index of the loop is the starting position of the current rune, measured in bytes, 
		and the code point is its value. 
	*/
	fmt.Printf("\n\n")
	const nihongo = "日本語"
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
}