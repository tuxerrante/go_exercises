// https://github.com/bitfield/tpg-tools2/tree/main/count
package counter

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
)

type counter struct {
	id      uuid.UUID
	input   io.Reader
	counted int
}

func NewCounter() *counter {
	return &counter{
		id:    uuid.New(),
		input: os.Stdin,
	}
}

func (c *counter) Lines() (int, error) {
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		log.Println(scanner.Text())
		c.counted++
	}
	log.Printf("Read %d lines", c.counted)
	return c.counted, nil
}
