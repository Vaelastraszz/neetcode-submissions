func evalRPN(tokens []string) int {

	stack := newStack()

	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for _, token := range tokens {

		if !operators[token] {
			val, _ := strconv.Atoi(token)
			stack.Push(val)
			continue
		}

		b := stack.Pop()
		a := stack.Pop()

		switch token {
		case "+":
			stack.Push(a + b)
		case "-":
			stack.Push(a - b)
		case "*":
			stack.Push(a * b)
		case "/":
			stack.Push(a / b)
		}
	}

	return stack.Top()
}

func parseNumber(s string) (int, bool) {
    n, err := strconv.Atoi(s)
    return n, err == nil
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
		return 0
	}

	last := len(s.data) - 1
	val := s.data[last]
	s.data = s.data[:last]

	return val
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	return s.data[len(s.data)-1]
}