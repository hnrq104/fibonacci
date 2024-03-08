package fibonacci

// adds the node x to the left of it's destination
func addNodeToCircularList(x, dst *FibonnaciNode) {
	//keep structure of underlying list
	x.left.right = x.right
	x.right.left = x.left

	x.right = dst
	x.left = dst.left

	x.right.left = x
	x.left.right = x
}

// Node links subjugates a node to another, returnin the one with lesser key
func nodeLink(x, y *FibonnaciNode) *FibonnaciNode {
	if x == nil || y == nil {
		panic("nodeLink called with nil arguments")
	}
	if y.key < x.key {
		x, y = y, x
	}
	y.parent = x

	y.left.right = y.right
	y.right.left = y.left

	if x.child != nil {
		y.right = x.child
		y.left = x.child.left
	} else {
		x.child = y
	}

	y.left.right = y
	y.right.left = y

	x.degree++
	y.mark = false

	return x
}

func (h *FibonnaciHeap) cut(x *FibonnaciNode) {
	if x.right == x {
		x.parent.child = nil
	} else if x.parent.child == x {
		x.parent.child = x.right
	}
	x.parent = nil
	addNodeToCircularList(x, h.min)
	x.mark = false
}

func (h *FibonnaciHeap) cascadingcut(x *FibonnaciNode) {
	z := x.parent
	if z == nil {
		return
	}
	if !x.mark {
		x.mark = true
	} else {
		h.cut(x)
		h.cascadingcut(z)
	}
}

func (h *FibonnaciHeap) lengthAndMaxDegree() (int, int) {
	if h.min == nil {
		return 0, 0
	}

	var length, maxdegree int

	aux := h.min
	for i := 0; i < 10; i++ {
		length++
		maxdegree = max(maxdegree, aux.degree)
		if aux.right == h.min {
			break
		}
		aux = aux.right
	}

	return length, maxdegree
}
