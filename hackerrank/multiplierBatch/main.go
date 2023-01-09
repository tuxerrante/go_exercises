package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 implement a server that generates a multiplication table a given integer by adding the value to the previous value
 in the sequence starting from 0.
 Given an integer toAdd the ith term of the sequence is given by (i-1)* toAdd where i>0
 the sequence for toAdd = 4 --> 0,4,8,12,16
 The server should send the results in batches of a given size. It should limit the rate of batches such that every batch is sent with
 delay interval of at least 10 ms
 the main function accepts 4 args
 - skipBatches = the num of batches to skip from the beginning
 - printBatches = the num of batches to be printed to the console
 - batchSize = the size of each batch
 - toAdd = the number for which the multiplication table is to be generated

 The server is expected to take 4 args
 - a bool channel that will accepts requests
 - the result int chan through which each batch will be sent
 - the size of the batch
 - the value for which the table is to be generated

 skipBatches = 0
 printBatches = 4
 batchSize = 3
 toAdd = 5
 --> print 4 batches of 3 numbers each. [0,5,10],[15,20,25][30,35,40],[45,50,55]
*/

/*
 * Complete the 'BurstyRateLimiter' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 *  3. int batchSize
 *  4. int toAdd
 */

var (
	lastTimestamp time.Time = time.Now()
	index         int       = 0
)

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {

	for <-requestChan {
		// Compute an entire batch as result
		for i := 0; i < batchSize; i++ {
			for time.Since(lastTimestamp) < time.Millisecond*10 {
				continue
			}
			lastTimestamp = time.Now()
			resultChan <- index * toAdd
			index += 1
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	skipBatches := int(skipTemp)

	totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	printBatches := int(totalTemp)

	batchSizeTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	batchSize := int(batchSizeTemp)

	toAddTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	toAdd := int(toAddTemp)

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go BurstyRateLimiter(requestChan, resultChan, batchSize, toAdd)

	for i := 0; i < skipBatches+printBatches; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		for j := 0; j < batchSize; j++ {
			new := <-resultChan
			if i < skipBatches {
				continue
			}
			fmt.Println(new)
		}
		if i >= skipBatches && i != skipBatches+printBatches-1 {
			fmt.Println("-----")
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff < 3 {
			fmt.Println("Rate is too high")
			break
		}
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
