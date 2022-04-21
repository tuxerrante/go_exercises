package main

import (
	"testing"
	"unicode/utf8"
)

var (
	str = "Hello, 世界"
	s   string
)

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func BenchmarkTrimFirstRune(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = trimFirstRune(str)
	}
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func BenchmarkTrimChar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s = trimLeftChar(str)
	}
}

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func BenchmarkTrimChars(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s = trimLeftChars(str, 1)
	}
}
