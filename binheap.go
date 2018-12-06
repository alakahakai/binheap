package binheap

import (
	"errors"
	"fmt"
)

type BinaryIntHeap struct {
	root  *BinaryIntHeapNode
	Count int
	IsMax bool
}

type BinaryIntHeapNode struct {
	Value      int
	Parent     *BinaryIntHeapNode
	LeftChild  *BinaryIntHeapNode
	RightChild *BinaryIntHeapNode
}

func (h *BinaryIntHeap) Insert(v int) int {
	var val int
	if h.IsMax {
		val = -v
	} else {
		val = v
	}
	if h.root == nil {
		h.root = new(BinaryIntHeapNode)
		h.root.Value = val
		h.Count += 1
		return v
	}
	// Find the right spot
	ptr := h.root
	for ptr != nil {
		if ptr.LeftChild == nil {
			newNode := new(BinaryIntHeapNode)
			newNode.Parent = ptr
			ptr.LeftChild = newNode
			ptr = newNode
			break
		} else if ptr.RightChild == nil {
			newNode := new(BinaryIntHeapNode)
			newNode.Parent = ptr
			ptr.RightChild = newNode
			ptr = newNode
			break
		} else if ptr.RightChild.Value < ptr.LeftChild.Value {
			ptr = ptr.RightChild
		} else {
			ptr = ptr.LeftChild
		}
	}
	ptr.Value = val
	// Bubble up
	for {
		if ptr.Parent == nil {
			break
		}
		if ptr.Parent.Value > ptr.Value {
			ptr.Value, ptr.Parent.Value = ptr.Parent.Value, ptr.Value
		}
		ptr = ptr.Parent
	}
	h.Count += 1
	return v
}

func (h *BinaryIntHeap) Pop() (*int, error) {
	if h.Count == 0 {
		return nil, errors.New("Heap is empty!")
	}
	if h.root == nil {
		panic(fmt.Sprintf("Heap root is nil, but Count is: %d", h.Count))
	}
	var v int
	if h.IsMax {
		v = -h.root.Value
	} else {
		v = h.root.Value
	}
	h.Count -= 1
	if h.Count == 0 {
		if h.root.LeftChild != nil || h.root.RightChild != nil {
			panic("Count is 0, but root has children!")
		}
		h.root = nil
		return &v, nil
	}
	// Find last child
	ptr := h.root
	for ptr != nil {
		if ptr.LeftChild == nil && ptr.RightChild == nil {
			break
		} else if ptr.LeftChild != nil && ptr.RightChild == nil {
			ptr = ptr.LeftChild
		} else if ptr.LeftChild == nil && ptr.RightChild != nil {
			ptr = ptr.RightChild
		} else if ptr.LeftChild.Value < ptr.RightChild.Value {
			if h.IsMax {
				ptr = ptr.RightChild
			} else {
				ptr = ptr.LeftChild
			}
		} else {
			if h.IsMax {
				ptr = ptr.LeftChild
			} else {
				ptr = ptr.RightChild
			}
		}
	}
	// Replace root with its value, and update its Parent
	h.root.Value = ptr.Value
	if ptr.Parent.LeftChild == ptr {
		ptr.Parent.LeftChild = nil
	} else if ptr.Parent.RightChild == ptr {
		ptr.Parent.RightChild = nil
	} else {
		panic("Parent pointer is not correct!")
	}
	// Bubble down
	ptr = h.root
	for ptr != nil {
		if ptr.LeftChild == nil && ptr.RightChild == nil {
			break
		} else if ptr.LeftChild != nil && ptr.RightChild == nil {
			if ptr.Value > ptr.LeftChild.Value {
				ptr.Value, ptr.LeftChild.Value = ptr.LeftChild.Value, ptr.Value
				ptr = ptr.LeftChild
			} else {
				break
			}
		} else if ptr.LeftChild == nil && ptr.RightChild != nil {
			if ptr.Value > ptr.RightChild.Value {
				ptr.Value, ptr.RightChild.Value = ptr.RightChild.Value, ptr.Value
				ptr = ptr.RightChild
			} else {
				break
			}
		} else {
			var smaller *BinaryIntHeapNode
			if ptr.LeftChild.Value < ptr.RightChild.Value {
				smaller = ptr.LeftChild
			} else {
				smaller = ptr.RightChild
			}
			if ptr.Value > smaller.Value {
				ptr.Value, smaller.Value = smaller.Value, ptr.Value
				ptr = smaller
			} else {
				break
			}
		}
	}
	return &v, nil
}

func (h *BinaryIntHeap) PopAll() []int {
	r := make([]int, 0)
	for {
		p, err := h.Pop()
		if err == nil {
			r = append(r, *p)
		} else {
			break
		}
	}
	return r
}

func NewBinaryIntHeap(data []int, isMax bool) *BinaryIntHeap {
	h := BinaryIntHeap{
		root:  nil,
		Count: 0,
		IsMax: isMax,
	}
	// Process data
	for _, v := range data {
		h.Insert(v)
	}
	return &h
}
