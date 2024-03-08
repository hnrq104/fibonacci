package fibonacci

import "log"

// Start with float64, then we see whatelse
type FibonnaciHeap struct {
	size  int
	roots int
	min   *FibonnaciNode
}

type FibonnaciNode struct {
	key    float64        // Key in the node
	degree int            // Number of children
	mark   bool           // Checks if a child was removed
	parent *FibonnaciNode // Parent of the node
	right  *FibonnaciNode // Sibling node to the right
	left   *FibonnaciNode // Sibling node to the left
	child  *FibonnaciNode // First child in circular linked list
}

func (h *FibonnaciHeap) Size() int {
	return h.size
}

func (x *FibonnaciNode) Key() float64 {
	return x.key
}

func NewFibonnaciHeap() FibonnaciHeap {
	return FibonnaciHeap{size: 0, min: nil, roots: 0}
}

func NewFibonnaciNode(k float64) *FibonnaciNode {
	n := &FibonnaciNode{key: k}
	n.left, n.right = n, n
	return n
}

// Inserts it as a new Tree in the Root List!
func (heap *FibonnaciHeap) InsertNode(x *FibonnaciNode) {
	x.degree = 0
	x.child = nil
	x.parent = nil
	x.mark = false

	if heap.min == nil {
		heap.min = x
	} else {
		x.right = heap.min
		x.left = heap.min.left
		x.left.right = x
		x.right.left = x

		if x.key < heap.min.key {
			heap.min = x
		}
	}
	heap.size++
	heap.roots++
}

func FibonnaciUnion(lhsHeap, rhsHeap FibonnaciHeap) FibonnaciHeap {
	if lhsHeap.min == nil {
		return rhsHeap
	}
	if rhsHeap.min == nil {
		return lhsHeap
	}

	heap := NewFibonnaciHeap()
	//concatenate both
	h1, h2 := lhsHeap.min, rhsHeap.min
	h1.left.right = h2.right
	h2.right.left = h1.left
	h2.right = h1
	h1.left = h2

	heap.min = h1
	if h2.key < heap.min.key {
		heap.min = h2
	}
	heap.size = lhsHeap.size + rhsHeap.size
	heap.roots = lhsHeap.roots + rhsHeap.roots
	return heap
}

func (heap *FibonnaciHeap) ExtractMin() *FibonnaciNode {
	removed := heap.min
	if removed == nil {
		return nil
	}

	nodeptr := removed.child
	// Worst do-while in my life
	if nodeptr != nil {
		for {
			aux := nodeptr
			nodeptr = nodeptr.right
			addNodeToCircularList(aux, heap.min)
			aux.parent = nil

			if aux == nodeptr {
				break
			}

		}
	}

	log.Println("passed first loop")
	//remove from root list
	removed.left.right = removed.right
	removed.right.left = removed.left

	if removed.right == removed {
		heap.min = nil
	} else {
		heap.min = removed.right
		heap.consolidate()
	}

	log.Println("passed consolidate")

	heap.size--
	heap.roots += removed.degree - 1
	return removed
}

func (h *FibonnaciHeap) consolidate() {
	rootlength, _ := h.lengthAndMaxDegree()
	newroots := make([]*FibonnaciNode, rootlength+1)
	aux := h.min
	for i := 0; i < rootlength; i++ {
		d := aux.degree
		for newroots[d] != nil {
			aux = nodeLink(aux, newroots[d])
			newroots[d] = nil
			d++
		}
		newroots[d] = aux
		if aux == aux.right {
			break
		}

		//remove aux from list
		aux.left.right = aux.right
		aux.right.left = aux.left

		aux = aux.right
	}

	for i := 0; i < len(newroots); i++ {
		if newroots[i] == nil {
			continue
		}

		if h.min == nil {
			h.min = newroots[i]
			h.min.left = h.min
			h.min.right = h.min
		} else {
			addNodeToCircularList(newroots[i], h.min)
			if newroots[i].key < h.min.key {
				h.min = newroots[i]
			}
		}
	}

}

func (h *FibonnaciHeap) DecreaseKey(x *FibonnaciNode, k float64) {
	if k > x.key {
		panic("Increasing Key value")
	}
	x.key = k
	y := x.parent

	if y != nil && x.key < y.key {
		h.cut(x)
		h.cascadingcut(y)
	}

	if h.min.key < x.key {
		h.min = x
	}
}
