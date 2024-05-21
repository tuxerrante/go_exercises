package counter

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBufferCounter(t *testing.T) {
	t.Parallel()
	t.Run("0_lines_FromStdin", func(t *testing.T) {
		c := NewBufferedCounter()
		c.input = *bufio.NewReader(strings.NewReader(""))
		got, err := c.Lines()
		if err != nil {
			t.Fatal(err)
		}
		exp := 0
		require.Equal(t, exp, got)
	})
	t.Run("1_lines_FromStdin", func(t *testing.T) {
		c := NewBufferedCounter()
		c.input = *bufio.NewReader(strings.NewReader("asd"))
		got, err := c.Lines()
		if err != nil {
			t.Fatal(err)
		}
		exp := 1
		require.Equal(t, exp, got)
	})
	t.Run("3_lines_FromStdin", func(t *testing.T) {
		c := NewBufferedCounter()
		c.input = *bufio.NewReader(strings.NewReader("a\nsd\n\n"))
		got, err := c.Lines()
		if err != nil {
			t.Fatal(err)
		}
		exp := 3
		require.Equal(t, exp, got)
	})
}

func TestCounter(t *testing.T) {
	t.Parallel()
	t.Run("0_lines_FromStdin", func(t *testing.T) {
		c := NewCounter()
		c.input = bytes.NewBufferString("")
		got, err := c.Lines()
		if err != nil {
			t.Fatal(err)
		}
		exp := 0
		require.Equal(t, exp, got)
	})
	t.Run("1_lines_FromStdin", func(t *testing.T) {
		c := NewCounter()
		c.input = bytes.NewBufferString("asd")
		got, err := c.Lines()
		if err != nil {
			t.Fatal(err)
		}
		exp := 1
		require.Equal(t, exp, got)
	})
	t.Run("3_lines_FromStdin", func(t *testing.T) {
		c := NewCounter()
		c.input = bytes.NewBufferString("a\nsd\n\n")
		got, err := c.Lines()
		if err != nil {
			t.Fatal(err)
		}
		exp := 3
		require.Equal(t, exp, got)
	})
}
