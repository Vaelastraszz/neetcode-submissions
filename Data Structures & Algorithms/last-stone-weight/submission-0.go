func lastStoneWeight(stones []int) int {

	if len(stones) == 0 {
		return 0
	}

	if len(stones) == 1 {
		return stones[0]
	}

	maxheap := &MaxHeap{}
	var stone1, stone2 int

	for _, stone := range stones {	
		heap.Push(maxheap, stone)	
	}

	for maxheap.Len() >= 2 {
		
		stone1 = heap.Pop(maxheap).(int)
		stone2 = heap.Pop(maxheap).(int)

		if stone1 != stone2 {
			heap.Push(maxheap, stone1 - stone2)
		}
	}

	if maxheap.Len() == 1 {
		return heap.Pop(maxheap).(int)
	}

	return 0

}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {

	*h = append(*h, x.(int))

}

func (h *MaxHeap) Pop() any {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x

}