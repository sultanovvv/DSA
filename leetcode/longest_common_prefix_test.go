package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	strs1 := []string{"qwer", "qw", "qwrt"}
	res1 := "qw"

	if res1 != longestCommonPrefix(strs1) {
		t.FailNow()
	}
}
