package container

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	// chang '<' to '>' then get max-heap
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func RunHeap() {
	h := &IntHeap{}
	heap.Init(h)

	nums := []int{2, 1, 5, 3, 4}
	fmt.Printf("* test min-heap with init int list: %+v\n", nums)

	for _, n := range nums {
		heap.Push(h, n)
	}

	fmt.Printf("heap sort: %+v\n", h)

	for len(*h) > 0 {
		x := heap.Pop(h)
		fmt.Printf("pop: %+v, heap after pop: %+v\n", x, h)
	}
}
