package leetcode

import (
	"math"
	"testing"
)

func TestMySqrt(t *testing.T) {

	x := 36

	res := mySqrt(x)
	if math.Floor(math.Sqrt(float64(x))) != float64(res) {
		t.FailNow()
	}

}
