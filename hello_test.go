package main

import (
	"testing"
)

func TestAddTwo(t *testing.T) {
	if AddTwo(3) != 5 {
		t.Error("expected 3 + 2 to equal 5")
	}
}

func TestTableAddTwo(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{2, 4},
		{0, 2},
		{-2, 0},
		{21, 23},
		{-5, -3},
		{-1, 1},
	}

	for _, test := range tests {
		if output := AddTwo(test.input); output != test.expected {
			t.Error("Test TableAddTwo failed: {} inputted, {} expected, but received: {}", test.input, test.expected, output)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		input1    string
		input2    string
		expected1 string
		expected2 string
	}{
		{"a", "b", "b", "a"},
		{"ab", "ba", "ba", "ab"},
		{"a", "a", "a", "a"},
	}

	for _, test := range tests {
		if output1, output2 := Swap(test.input1, test.input2); output1 != test.expected1 && output2 != test.expected2 {
			t.Error("Test TableAddTwo failed: {} and {} inputted, {} and {} expected, but received: {} and {}", test.input1, test.input2, test.expected1, test.expected1, output1, output2)
		}
	}
}
