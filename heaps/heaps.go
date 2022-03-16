package heaps

import "fmt"

type MaxHeap struct {
	array []int
}

func Make_Heap() MaxHeap {
    var myheap MaxHeap = MaxHeap{}
    return myheap
}

func (heap *MaxHeap) Insert(data int) {
	heap.array = append(heap.array,data)
	heap.add(len(heap.array)-1)
}

func (heap *MaxHeap) add(index int) {
	for heap.array[parent(index)] < heap.array[index] {
		heap.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i-1)/2
}

func left(i int) int {
	return 2*i+1
}
func right(i int) int {
	return 2*1+2
}

func (heap *MaxHeap) swap(i1, i2 int) {
	heap.array[i1], heap.array[i2] = heap.array[i2], heap.array[i1]
}

func (heap *MaxHeap) Remove() int {
	extracted := heap.array[0]
	l := len(heap.array)-1
	if l == 0 {
		fmt.Println("can't extract, len = 0")
		return -1
	}
	heap.array[0] = heap.array[l]
	heap.array = heap.array[:l]
	heap.restore(0)
	return extracted
}

func (heap *MaxHeap) restore(index int) {
	lastIndex := len(heap.array)-1
	l,r := left(index), right(index)
	compare := 0
	for l <= lastIndex {
		if l == lastIndex {
			compare = l
		} else if heap.array[l] > heap.array[r] {
			compare = l
		} else {
			compare = r
		}

		if heap.array[index] < heap.array[compare] {
			heap.swap(index, compare)
			index = compare
			l,r = left(index), right(index)
		} else {
			return
		}
	} 
}

func (heap *MaxHeap) Print_Heap() {
	for _, v := range heap.array {
		fmt.Print(v," ")
	}
	fmt.Println()
}
