package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "sort"
    "strconv"
    "strings"
)
/*
 * Complete the 'RemainderSorting' function below (and other code for sorting if needed).
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.

	REFERENCES:
	https://yourbasic.org/golang/how-to-sort-in-go/
	https://pkg.go.dev/sort#Interface
 */


func RemainderSorting(strArr []string) []string {

    type modStrings struct{
        mod     int
        value   string
    }
    myCustomStrings := []modStrings{}
    
    for _, s := range strArr{
        curentModString := modStrings { 
            mod:    len(s)%3,
            value:  s,
        }
        myCustomStrings = append(myCustomStrings, curentModString)
        fmt.Println(myCustomStrings)
    }

    // sort by modulo 3
    sort.Slice(myCustomStrings, func(i, j int) bool { return myCustomStrings[i].mod < myCustomStrings[j].mod })
    fmt.Println(myCustomStrings)

    // sort again only strings with same module, alphabetically
    sort.SliceIsSorted(myCustomStrings, func(i, j int) bool {return myCustomStrings[i].value < myCustomStrings[j].value })
    
    returnSlice := []string{}
    for _,v := range myCustomStrings {
        returnSlice = append(returnSlice, v.value)
    }
    return returnSlice
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var strArr []string

    for i := 0; i < int(strArrCount); i++ {
        strArrItem := readLine(reader)
        strArr = append(strArr, strArrItem)
    }

    result := RemainderSorting(strArr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
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
