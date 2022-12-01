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
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	readFile, err := os.Open("input1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	h := &IntHeap{}
	heap.Init(h)

	var sum int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			heap.Push(h, sum)
			sum = 0
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += val
		}
	}
	var top1 int = heap.Pop(h).(int)
	var top2 int = heap.Pop(h).(int)
	var top3 int = heap.Pop(h).(int)
	fmt.Println("Day 1 task 1: ", top1)
	fmt.Println("Day 1 task 2: ", top1+top2+top3)
}
