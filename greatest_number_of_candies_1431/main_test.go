package gnc1431

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies []int
		extra   int
		expect  []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, test := range tests {
		result := kidsWithCandies(test.candies, test.extra)
		if !reflect.DeepEqual(result, test.expect) {
			t.Errorf("mergeAlternately failed: expected  %v but got %v", test.expect, result)
		}
	}
}
