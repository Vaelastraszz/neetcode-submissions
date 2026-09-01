func findKthLargest(nums []int, k int) int {
	h := MaxHeap(nums)
	heap.Init(&h)

	var val int

	for i:= 1; i <= k; i++ {
		val = heap.Pop(&h).(int)
	}

	return val
}

type MaxHeap []int 

func(h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i,j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
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