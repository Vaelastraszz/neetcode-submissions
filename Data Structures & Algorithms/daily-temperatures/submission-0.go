func dailyTemperatures(temperatures []int) []int {

	s := newStack()
	result := make([]int, len(temperatures), len(temperatures))

	for i, temp := range temperatures {

		if s.Top() < 0 {
			s.Push(i)
			continue
		}

		if temp > temperatures[s.Top()] {

			for temperatures[s.Top()] < temp {
				pos := s.Pop()
				result[pos] = i - pos
				if s.Top() < 0 {
					break
				}
			}

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