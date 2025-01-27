package model

import (
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

func (heap PriorityQueue) Swap(i, j int) {
    heap[i], heap[j] = heap[j], heap[i]
}

func (heap *PriorityQueue) siftUp(index int) {
    parent := (index - 1) / 2
    for parent >= 0 {
        if ((*heap)[index].Freq > (*heap)[parent].Freq) {
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
    for left <= heapLen {
        swap := left
        if right <= heapLen && (*heap)[swap].Freq > (*heap)[right].Freq {
            swap = right
        }
        if (*heap)[index].Freq > (*heap)[swap].Freq {
			heap.Swap(index, swap)
			index = swap
		} else {
			break
		}
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
