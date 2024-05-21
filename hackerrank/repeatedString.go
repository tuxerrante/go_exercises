/*
There is a string, , of lowercase English letters that is repeated infinitely many times. Given an integer, , find and print the number of letter a's in the first letters of the infinite string.

Example
The substring we consider is , the first characters of the infinite string. There are occurrences of a in the substring.

repeatedString has the following parameter(s):

	s: a string to repeat
	n: the number of characters to consider

Returns

	int: the frequency of a in the substring

The first line contains a single string,
The second line contains an integer,

Constraints
1<= s <= 100
1<= n <= 10^12
*/
package hackerrank

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int64) int64 {
	// No need to count matches after the first s repetition
	// (n / len(s)) * count(s, "a") + count(s[ : n % len(s)] )
	var matches int64
	var stringLen int64 = int64(len(s))

	// Count matches in the main string
	if n > stringLen {
		matches = countByte(s, 'a')
	}

	// Count matches in the final substring
	extraMatches := countByte(s[:n%stringLen], 'a')

	// Multiply the matches for the string repetitions
	return n/stringLen*matches + extraMatches
}

func countByte(s string, b byte) int64 {
	var matches int64

	i := strings.IndexByte(s, 'a')
	if i < 0 {
		return 0
	}
	matches++
	s = s[i+1:]

	for len(s) > 0 {
		i := strings.IndexByte(s, 'a')
		if i < 0 {
			break
		}
		matches++
		s = s[i+1:]
	}
	return matches
}

func countBytesTest() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

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
