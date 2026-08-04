func dailyTemperatures(temperatures []int) []int {

	s := newStack()
	result := make([]int, len(temperatures))

	for i, temp := range temperatures {

		for s.Top() != -1 && temp > temperatures[s.Top()] {
			pos := s.Pop()
			result[pos] = i - pos
		}

		s.Push(i)
	}

	return result
}

type Stack struct {
	data []int
}

func newStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}

	last := len(s.data) - 1
	val := s.data[last]
	s.data = s.data[:last]

	return val
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return -1
	}

	return s.data[len(s.data)-1]
}