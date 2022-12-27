package model

type Dequeue struct {
	innerSlice []interface{}
}

func NewDequeue(s ...interface{}) *Dequeue {

	return &Dequeue{innerSlice: s}
}

func (d *Dequeue) Poll() interface{} {
	if len(d.innerSlice) == 0 {
		return nil
	}
	res := d.innerSlice[0]
	d.innerSlice = d.innerSlice[1:]
	return res
}

func (d *Dequeue) Offer(elem interface{}) bool {
	d.innerSlice = append(d.innerSlice, elem)
	return true
}

func (d *Dequeue) IsEmpty() bool {
	return len(d.innerSlice) == 0
}
