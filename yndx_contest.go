package main

import (
	"fmt"
)

func stonesAndTreasures() {

	var s, j string

	_ = fmt.Sprintf("%s", s)
	_ = fmt.Sprintf("%s", j)

	mSet := map[rune]struct{}{}

	for _, sstr := range s {
		mSet[sstr] = struct{}{}
	}

	res := 0

	for _, jstr := range j {
		if _, ok := mSet[jstr]; ok {
			res++
		}
	}

	fmt.Println(res)
}
