func kClosest(points [][]int, k int) [][]int {
	h := MinHeap(points)
	heap.Init(&h)

	res := [][]int{}

	for i:=1; i<=k; i++ {
		res = append(res, heap.Pop(&h).([]int))
	}

	return res
}

type MinHeap [][]int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {
    
    dist1 := h[i][0]*h[i][0] + h[i][1]*h[i][1]
    dist2 := h[j][0]*h[j][0] + h[j][1]*h[j][1]

	return dist1 < dist2
}

func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}