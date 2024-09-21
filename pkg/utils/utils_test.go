package utils

import "testing"

func TestRange(t *testing.T) {
	lines := []string{"1", "2", "3", "  4 ", " 5    "}
	for num, err := range ParseIntRange(lines) {
		if err != nil {
			t.Fatal(err)
		}
		t.Log(num)
	}
}
