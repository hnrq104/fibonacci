package fibonacci

import (
	"testing"
)

func TestInsertAndExtract(t *testing.T) {
	h := NewFibonnaciHeap()

	n := 1000
	var m int
	for m = 0; m < n; m++ {
		h.InsertNode(NewFibonnaciNode(float64(m)))
	}

	if h.Size() != n {
		t.Errorf("heap.Size() = %d, expect %d", h.Size(), n)
	}

	for m > 0 {
		h.ExtractMin()
		m--
	}

	if h.Size() != 1 {
		t.Errorf("heap.Size() = %d, expect %d", h.Size(), 1)
	}
}

func TestDecreaseKey(t *testing.T) {
	n := 1000
	fibnodes := make([]*FibonnaciNode, n)
	fh := NewFibonnaciHeap()
	for i := range fibnodes {
		fn := NewFibonnaciNode(float64(i))
		fibnodes[i] = fn
		fh.InsertNode(fn)
	}

	for i := 0; i < n; i++ {
		if i%10 == 0 {
			fh.ExtractMin()
			continue
		}
		fh.DecreaseKey(fibnodes[i], float64(i-1))
	}

	//Checks if panics
}
