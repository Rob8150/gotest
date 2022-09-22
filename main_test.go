package main

import "testing"

// :GoTest (in neovim)
// go test -v  (verbose)
func TestCalculator(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{9999, 10001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} expected, recieved: {}", test.input, test.expected, output)
		}
	}

}
