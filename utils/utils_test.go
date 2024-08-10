package utils

import (
	"testing"
)

func TestExtractNums(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{
			"12",
			[]int{12},
		},
		{
			"ab12arst3",
			[]int{12, 3},
		},
		{
			"ab12arst3s4arst",
			[]int{12, 3, 4},
		},
		{
			"absee",
			[]int{},
		},
	}

	for _, test := range tests {
		result := ExtractNums(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("expected %v, got %v.", test.expected, result)
		}
		for i, num := range result {
			if num != test.expected[i] {
				t.Errorf("expected %v, got %v.", test.expected, result)
			}
		}
	}
}
