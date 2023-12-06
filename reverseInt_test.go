package main

import "testing"

var tests_reverseInt = []struct {
	num      int
	expected int
}{
	{15, 51},
	{981, 189},
	{500, 5},
	{-15, -51},
	{-90, -9},
	{0, 0},
}

func TestReverseInt(t *testing.T) {
	for _, test := range tests_reverseInt {
		result := reverseInt(test.num)
		if result != test.expected {
			t.Errorf("Expected %v but got %v", test.expected, result)
		}
	}
}
