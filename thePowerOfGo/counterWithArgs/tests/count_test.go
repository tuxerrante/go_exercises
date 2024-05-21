package tests

import (
	"testing"
	count "github.com/tuxerrante/go_exercises/power-of-go/counterwithargs/args_count"
)

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()

	args := []string{"./testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}

	got := c.Lines()
	want := 3
	if want != got {
		// Prefer t.Error here instead of t.Fatal
		t.Errorf("Got %d Want %d", got, want)
	}
}
