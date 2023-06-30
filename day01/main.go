package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	file, err := os.Open("input.data")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	h := &IntHeap{}
	heap.Init(h)
	sum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			val, _ := strconv.Atoi(line)
			sum += val
		} else {
			heap.Push(h, sum)
			if len(*h) > 3 {
				heap.Pop(h)
			}
			sum = 0
		}
	}

	top := 0
	i := 3
	for len(*h) > 0 {
		val := heap.Pop(h).(int)
		fmt.Printf("Best %d has %d\n", i, val)
		top += val
		i--
	}

	fmt.Printf("Total: %d\n", top)
	file.Close()
}
