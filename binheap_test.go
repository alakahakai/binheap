package binheap

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func reverseSlice(a []int) {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}

func isSameSlice(a, b []int) bool {
	i := len(a)
	j := len(b)
	if i != j {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func TestBinMaxHeap(t *testing.T) {
	arr := []int{}
	isMax := true
	h := NewBinaryIntHeap(arr, isMax)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	j := r.Intn(1000)
	for i := 0; i < j; i++ {
		m := r.Int()
		n := r.Intn(2)
		if n == 0 {
			h.Insert(-m)
			arr = append(arr, -m)
		} else {
			h.Insert(m)
			arr = append(arr, m)
		}
	}
	d := h.PopAll()
	sort.Ints(arr)
	reverseSlice(arr)
	if !isSameSlice(d, arr) {
		t.Errorf("Test failed!")
		t.Errorf("Binheap sorted: %v", d)
		t.Errorf("Golang library sorted: %v", arr)
	}
}

func TestBinMinHeap(t *testing.T) {
	arr := []int{}
	isMax := false
	h := NewBinaryIntHeap(arr, isMax)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	j := r.Intn(10000)
	for i := 0; i < j; i++ {
		m := r.Int()
		n := r.Intn(2)
		if n == 0 {
			h.Insert(-m)
			arr = append(arr, -m)
		} else {
			h.Insert(m)
			arr = append(arr, m)
		}
	}
	d := h.PopAll()
	sort.Ints(arr)
	if !isSameSlice(d, arr) {
		t.Errorf("Test failed!")
		t.Errorf("Binheap sorted: %v", d)
		t.Errorf("Golang library sorted: %v", arr)
	}
}

func BenchmarkBinHeap(b *testing.B) {
	isMax := true
	h := NewBinaryIntHeap([]int{}, isMax)
  b.ResetTimer()
	for i := 0; i < b.N; i++ {
			h.Insert(i)
	}
}
