package main

import (
	"testing"
)

var tests_maxChars = []struct {
	str     string
	maxChar string
	count   int
}{
	{"abccccccccd", "c", 8},
	{"aaa bjhbjdfbvldkj 7777777", "7", 7},
}

func TestMaxChar(t *testing.T) {
	for _, test := range tests_maxChars {
		maxChar, count := maxChar(test.str)
		if test.maxChar != maxChar || test.count != count {
			t.Errorf("Expected %v, %v but got %v, %v", test.maxChar, test.count, maxChar, count)
		}
	}
}
