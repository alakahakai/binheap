package main

import (
	"fmt"
	"github.com/rayqiu/binheap"
	"math/rand"
	"sort"
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

func main() {
	arr := []int{}
	isMax := true
	h := binheap.NewBinaryIntHeap(arr, isMax)
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
	if isSameSlice(d, arr) {
		fmt.Println("Test successful!")
		fmt.Println("Sorted:", arr)
	} else {
		fmt.Println("Test failed!")
		fmt.Println("Binheap sorted:", d)
		fmt.Println("Golang library sorted:", arr)
	}
	arr = []int{}
	isMax = false
	h = binheap.NewBinaryIntHeap(arr, isMax)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	j = r.Intn(10000)
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
	d = h.PopAll()
	sort.Ints(arr)
	if isSameSlice(d, arr) {
		fmt.Println("Test successful!")
		fmt.Println("Sorted:", arr)
	} else {
		fmt.Println("Test failed!")
		fmt.Println("Binheap sorted:", d)
		fmt.Println("Golang library sorted:", arr)
	}
}
