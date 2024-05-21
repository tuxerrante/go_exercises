// https://github.com/bitfield/tpg-tools2/tree/main/count
package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

type bufferedCounter struct {
	id      uuid.UUID
	input   bufio.Reader
	counted int
}

func NewBufferedCounter() *bufferedCounter {

	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error %s", err)
	}

	return &bufferedCounter{
		id:    uuid.New(),
		input: *bufio.NewReaderSize(os.Stdin, int(stat.Size())),
	}
}

func (c *bufferedCounter) Lines() (int, error) {

	if c.input.Size() == 0 {
		return 0, nil
	}

	for {
		b, err := c.input.ReadBytes('\n')
		if err == io.EOF || b == nil {
			if len(b) > 0 {
				c.counted++
			}
			break
		} else if err != nil {
			return 0, fmt.Errorf("error reading input bytes: %v", err)
		}
		c.counted++
	}

	return c.counted, nil
}
