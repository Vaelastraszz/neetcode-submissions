type MinStack struct {
	data    []int
	allMins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)

	if len(s.allMins) == 0 || val <= s.allMins[len(s.allMins)-1] {
		s.allMins = append(s.allMins, val)
	}
}

func (s *MinStack) Pop() {
	if len(s.data) == 0 {
		return
	}

	last := len(s.data) - 1
	val := s.data[last]
	s.data = s.data[:last]

	if val == s.allMins[len(s.allMins)-1] {
		s.allMins = s.allMins[:len(s.allMins)-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.allMins) == 0 {
		return 0
	}
	return s.allMins[len(s.allMins)-1]
}