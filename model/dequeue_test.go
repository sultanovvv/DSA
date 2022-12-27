package model

import "testing"

func TestDequeue_Poll(t *testing.T) {

	elem1 := "test1"
	elem2 := "test2"
	elem3 := "test3"

	q := NewDequeue()
	q.Offer(&elem1)
	q.Offer(&elem2)
	q.Offer(&elem3)

	if len(q.innerSlice) != 3 {
		t.FailNow()
	}

	ep := q.Poll().(*string)

	if ep != &elem1 {
		t.FailNow()
	}

	if len(q.innerSlice) != 2 {
		t.FailNow()
	}

}
