func isValid(s string) bool {
    
	if len(s) == 0 || len(s) % 2 != 0 {
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

		if stack.Peek() != 0 {
			last_par := stack.Pop()
			if matching_pair, ok := matchingMap[last_par]; ok {
				if c == matching_pair {
					continue
				}

				return false
			}
		}

		return false
	}

	return true && stack.Peek() == 0
}

type Stack struct {
	data []rune
}

func newStack() *Stack {
	
	s:= make([]rune, 0)
	
	return &Stack{
		data : s,
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

	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return res
}