package model

type Queue struct {
	innerSlice []interface{}
}

func NewQueue(s ...interface{}) *Queue {

	return &Queue{innerSlice: s}
}

func (d *Queue) Poll() interface{} {
	if len(d.innerSlice) == 0 {
		return nil
	}
	res := d.innerSlice[0]
	d.innerSlice = d.innerSlice[1:]
	return res
}

func (d *Queue) Offer(elem interface{}) bool {
	d.innerSlice = append(d.innerSlice, elem)
	return true
}

func (d *Queue) IsEmpty() bool {
	return len(d.innerSlice) == 0
}
