package calculationutil

import "testing"

func TestAddTwoNum(t *testing.T) {
	if got := AddTwoNum(2, 4); got != 6 {
		t.Fatalf("\"add testing failed\": expected %v  but got %v\n", 6, got)
	}
}


func TestSubBFromA(t *testing.T) {
	if got := SubBFromA(6, 4); got != 2 {
		t.Fatalf("\"SubBFromA failed\": expected %v  but got %v\n", 2, got)
	}
}