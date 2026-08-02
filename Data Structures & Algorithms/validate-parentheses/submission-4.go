func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	matchingMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := newStack()

	for _, c := range s {
		if _, ok := matchingMap[c]; ok {
			stack.Push(c)
			continue
		}

		if stack.Peek() == 0 {
			return false
		}

		last := stack.Pop()
		if c != matchingMap[last] {
			return false
		}
	}

	return stack.Peek() == 0
}

type Stack struct {
	data []rune
}

func newStack() *Stack {
	return &Stack{
		data: make([]rune, 0),
	}
}

func (s *Stack) Push(c rune) {
	s.data = append(s.data, c)
}

func (s *Stack) Peek() rune {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() rune {
	if len(s.data) == 0 {
		return 0
	}

	last := len(s.data) - 1
	res := s.data[last]
	s.data = s.data[:last]

	return res
}