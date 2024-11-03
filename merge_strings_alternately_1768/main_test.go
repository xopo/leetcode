package merge

import (
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		word1, word2, expected string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, test := range tests {
		result := mergeAlternately(test.word1, test.word2)
		if result != test.expected {
			t.Errorf("mergeAlternately failed: expected %q but got %q", test.expected, result)
		}
	}
}
