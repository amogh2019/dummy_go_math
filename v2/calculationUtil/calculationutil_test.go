package calculationutil

import "testing"

func TestAddTwoNum(t *testing.T) {
	if got := AddTwoNum(2, 4); got != 6 {
		t.Fatalf("\"add testing failed\": case %v  but got %v\n", 6, got)
	}
}
