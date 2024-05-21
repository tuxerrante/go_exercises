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
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// keep the num of walked valleys
	var valleyCounter int32

	// Start walking the path
	for i := int32(0); i < steps; i++ {
		if rune(path[i]) == 'D' {
			// A valley is a sequence of **consecutive** steps below sea level
			fmt.Printf("Going down in the valley!\n D(-1)")
			i = walk(i+1, path, -1, &valleyCounter)

		} else if rune(path[i]) == 'U' {
			fmt.Printf("Going UP on the mountain!\n U(1)")
			i = walk(i+1, path, 1, &valleyCounter)
		}
	}
	return valleyCounter
}

// ASCII characters are single-byte corresponding to the first 128 Unicode characters.
//
//	Strings behave like slices of bytes. A rune is an integer value identifying a Unicode code point.
func walk(currentPosition int32, path string, seaLevel int, valleyCounter *int32) int32 {
	pathLen := int32(len(path))

	// Check for monotonic descent
	var downSteps, upSteps int32
	var couldBeAValley bool
	if seaLevel < 0 {
		downSteps += 1
		couldBeAValley = true
	} else {
		upSteps += 1
	}

	for ; seaLevel != 0 && currentPosition < pathLen; currentPosition += 1 {

		currentRune := rune(path[currentPosition])

		if currentRune == 'D' {
			seaLevel -= 1
			downSteps += 1
			fmt.Printf("D(%d)", seaLevel)

		} else if currentRune == 'U' {
			seaLevel += 1
			upSteps += 1
			fmt.Printf("U(%d)", seaLevel)
		}
	}
	if seaLevel == 0 {
		fmt.Printf("\n_ We're back on sea level.\n")
	}

	if downSteps == upSteps && couldBeAValley {
		*valleyCounter += 1
	}
	return currentPosition
}

func walkTest() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	valleyCheckError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	valleyCheckError(err)
	steps := int32(stepsTemp)

	path := valleyReadLine(reader)

	result := countingValleys(steps, path)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func valleyReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func valleyCheckError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
    8
    UDDDUDUU

    Going UP on the mountain!
     U(1)D(0)
    _ We're back on sea level.
    Going down in the valley!
     D(-1)U(0)
    _ We're back on sea level.
    Going UP on the mountain!
     U(1)U(2)
-----
    12
    DDUUDDUDUUUD

    Going down in the valley!
     D(-1)D(-2)U(-1)U(0)
    _ We're back on sea level.
    Going down in the valley!
     D(-1)U(0)
    _ We're back on sea level.
    Going UP on the mountain!
     U(1)U(2)U(3)D(2)
*/
