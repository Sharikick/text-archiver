package model

import (
	"fmt"
	"log/slog"
	"os"
)

type Node struct {
    Left, Right *Node
    Freq        int
    Char        rune
}

type PriorityQueue []*Node

func (heap PriorityQueue) Len() int {
    return len(heap)
}

func (heap PriorityQueue) Less(i, j int) bool {
    return heap[i].Freq > heap[j].Freq
}

func (heap PriorityQueue) Swap(i, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (heap *PriorityQueue) siftUp(index int) {
    parent := (index - 1) / 2
    for parent >= 0 {
        if heap.Less(parent, index) {
            heap.Swap(index, parent)
        } else {
            break
        }

        index = parent
        parent = (index - 1) / 2
    }
}

func (heap *PriorityQueue) siftDown(index int) {
    left := 2*index+1
    right := 2*index+2
    heapLen := heap.Len()
    for left < heapLen {
        swap := left
        if right < heapLen && heap.Less(swap, right) {
            swap = right
        }
        if heap.Less(index, swap) {
			heap.Swap(index, swap)
			index = swap
		} else {
			break
		}
    }
}

func (heap PriorityQueue) Print() {
    for _, node := range heap {
        fmt.Println(*node)
    }
}

func (heap *PriorityQueue) Add(node *Node) {
	*heap = append(*heap, node)
	heap.siftUp(heap.Len() - 1)
}

func (heap *PriorityQueue) Pop() *Node {
    if heap.Len() == 0 {
		slog.Error("The list is empty")
		os.Exit(1)
	}

	node := (*heap)[0]
	lastIndex := heap.Len() - 1
	(*heap)[0] = (*heap)[lastIndex]
	*heap = (*heap)[:lastIndex]
	heap.siftDown(0)
	return node
}
