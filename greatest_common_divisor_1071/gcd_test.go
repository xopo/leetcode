package gcd1071

import (
	"testing"
)

func TestGcd(t *testing.T) {
	tests := []struct {
		i []string
		o string
	}{
		{i: []string{"ABCABC", "ABC"}, o: "ABC"},
		{i: []string{"ABABAB", "ABAB"}, o: "AB"},
		{i: []string{"LEET", "CODE"}, o: ""},
		{
			i: []string{
				"TAUXXTAUXXTAUXXTAUXXTAUXX",
				"TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			},
			o: "TAUXX",
		},
	}
	for _, test := range tests {
		result := gcdOfStrings(test.i[0], test.i[1])
		if result != test.o {
			t.Errorf("Expected %q but got %q", test.o, result)
		}
	}
}
