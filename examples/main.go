package main

import (
	"github.com/rayqiu/binheap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func isSameSlice(a, b []int, reversed bool) bool {
	i := len(a)
	j := len(b)
	if i != j {
		return false
	}
	for k, v := range a {
		if reversed {
			if v != b[j-k-1] {
				return false
			}
		} else {
			if v != b[k] {
				return false
			}
		}
	}
	return true
}

func main() {
	arr := []int{}
	isMax := true
	h := binheap2.NewBinaryIntHeap(arr, isMax)
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
	if isSameSlice(d, arr, isMax) {
		fmt.Println("Test successful!")
		fmt.Println("Binheap sorted:", d)
		fmt.Println("Golang library sorted:", arr)
	} else {
		fmt.Println("Test failed!")
		fmt.Println("Binheap sorted:", d)
		fmt.Println("Golang library sorted:", arr)
	}
}
