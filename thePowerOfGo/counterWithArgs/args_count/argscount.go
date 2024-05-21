package args_count

import (
	"bufio"
	"io"
	"os"
)

// No need to export the data structure
// since we make use of a constructor and setter funcs
type counter struct {
	input  io.Reader
	output io.Writer
}

// Defining object fields setters of the same type
// to have a type safe dynamic way to set defaults
// without the need to pass nil values
type counterOption func(*counter) error

// Apply custom values on a new counter object fields if provided
// otherwise the defaults will be set on stdin/stdout
func NewCounter(opts ...counterOption) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// This func returns an option func in order to configure the constructor
func WithInputFromArgs(args []string) counterOption {
	return func(c *counter) error {
		file, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = file
		return nil
	}
}

// Returns number of lines read from counter input
func (c *counter) Lines() int {
	var lines int
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	return lines
}
