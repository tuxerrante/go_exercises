package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverseString(f *testing.F) {

	testCases := []string{
		"", "a", "ab", "123!", "🐌⚠️", "泃",
	}

	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, tc string) {
		got, err := ReverseString(tc)
		t.Logf("test case: %s", tc)
		if err != nil {
			if utf8.ValidString(tc) && !utf8.ValidString(got) {
				t.Error()
			}
		}
	})
}
