package comparisonUtil

import (
	"strconv"
	"testing"
)

func TestCompareSliceScenariosIndividually(t *testing.T) {

	type testCase struct {
		slice1   []int
		slice2   []int
		expected bool
	}

	testCases := []testCase{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
			false,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3, 4},
			false,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2},
			false,
		},
		{
			[]int{1, 2, 3},
			nil,
			false,
		},
		{
			nil,
			nil,
			true,
		},
	}

	for i, ithCase := range testCases {
		t.Run("Compare Slice Test "+strconv.Itoa(i), func(t *testing.T) {
			if got := CompareSlices(ithCase.slice1, ithCase.slice2); got != ithCase.expected {
				t.Fatalf("\"slice comparison failed\": case %v  but got %v\n", ithCase, got)
			}
		})

	}
}
