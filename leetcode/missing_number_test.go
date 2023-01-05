package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	missing := 4
	n := missingNumber([]int{5, 0, 1, 6, 2, 3, 7})
	if n != missing {
		t.FailNow()
	}

	missing = 2
	n = missingNumber([]int{5, 0, 1, 6, 7, 3, 4})
	if n != missing {
		t.FailNow()
	}
}
