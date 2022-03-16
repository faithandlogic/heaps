package main

import (
	"heaps/heaps"
)

func main() {
	println("Heaps")

	heap := heaps.Make_Heap()
	buildHeap := []int{12,22,32,5,3,7,2,23,44,32,11}
	for _, v := range buildHeap {
		heap.Insert(v)
	}
	heap.Print_Heap()
	for i := 0; i < 5; i++ {
		heap.Remove()
	}
	heap.Print_Heap()

}