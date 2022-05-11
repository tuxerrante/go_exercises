package main
/*
	https://www.hackerrank.com/challenges/sparse-arrays/problem

	To make it work locally
	export OUTPUT_PATH="./out"
*/
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
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */
func matchingStrings(strings []string, queries []string) []int32 {
    
    matches := make(map[string]int32, len(queries))
    counter := make([]int32, len(queries))
    
    for _,q := range queries {
        matches[q] = 0
    }
    // use a map to count, remembering ordering is not garanteed
    for _,s := range strings {
        _,ok := matches[s]
        if ok {
            matches[s] ++
            fmt.Printf("Match on string %s, increasing to %v\n", s, matches[s])
        }
    }
    // append put the values on the slice tail, instead I want the same indexes
    for idx, q := range queries {
        counter[idx] = matches[q]
        fmt.Printf("%v\n", counter)
    }
    return counter
}

func main() {
    //reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	file , err := os.Open("test-sparseArrays")
	if err != nil {
		log.Fatal("Error opening input file!")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    stringsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var strings_slice []string

    for i := 0; i < int(stringsCount); i++ {
        stringsItem  := readLine(reader)
        strings_slice = append(strings_slice, stringsItem)
    }

    queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var queries []string

    for i := 0; i < int(queriesCount); i++ {
        queriesItem := readLine(reader)
        queries = append(queries, queriesItem)
    }

    res := matchingStrings(strings_slice, queries)

    for i, resItem := range res {
        fmt.Fprintf(writer, "%d", resItem)

        if i != len(res) - 1 {
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
