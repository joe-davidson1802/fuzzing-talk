package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) { // Tests are functions
	testcases := []struct { // Run multiple tests using slices
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in) // Act
		if rev != tc.want {   // Assert
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) { // Slightly different function name
	f.Add("Hello, world") // Seed the corpus!
	f.Add(" ")
	f.Add("!12345")
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func FuzzReverseTwice(f *testing.F) { // Slightly different function name
	f.Add("Hello, world") // Seed the corpus!
	f.Add(" ")
	f.Add("!12345")
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
