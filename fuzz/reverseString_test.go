package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		in, want string
	}{
		{
			"radar",
			"radar",
		},
		{
			"pippo",
			"oppip",
		},
	}
	for _, test := range testCases {
		got, _ := ReverseString(test.in)
		if got != test.want {
			t.Errorf("Reverse: %q, want %q", got, test.want)
		}
	}
}
