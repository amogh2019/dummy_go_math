package comparisonUtil

import (
	"fmt"
	"testing"
)

func TestMaxInt(t *testing.T) {

	got := MaxInt(10, 20)
	want := 20
	if got != want {
		t.Logf("\"max is not expected\": %v\n", "max is not expected")
		t.FailNow()
	}

	got = MaxInt(20, 10)
	if got != want {
		t.Errorf("\"max is not expected\": %v\n", "max is not expected") // errorf == logf + fail // fatalf = logf + failnow
	}

}

func TestMinIntScenarios(t *testing.T) {

	type testCase struct {
		num1     int
		num2     int
		expected int
	}

	testCases := []testCase{
		{num1: 1, num2: 2, expected: 1},
		{num1: 2, num2: 1, expected: 1},
	}

	for _, ithcase := range testCases {
		if got := MinInt(ithcase.num1, ithcase.num2); got != ithcase.expected {
			fmt.Printf("\"min is not expected\": %v\n", "min is not expected")
			t.FailNow()
		}
	}
}

func BenchmarkMinInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		MinInt(4, 6)
	}
}
