func leastInterval(tasks []byte, n int) int {
	freq := [26]int{}

	for _, task := range tasks {
		freq[task - 'A'] ++ 
	}

	h := MaxHeap(freq[:])
	heap.Init(&h)

	maxFreq := heap.Pop(&h).(int)
	nBlocs := maxFreq - 1

	nCycles := maxFreq + nBlocs*n

	for len(h) > 0 && h[0] == maxFreq {
		heap.Pop(&h)
		nCycles++
	}

	if len(tasks) > nCycles {
		return len(tasks)
	}

	return nCycles
}

type MaxHeap []int

func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h MaxHeap) Less(i, j int) bool {return h[i] > h[j]}
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