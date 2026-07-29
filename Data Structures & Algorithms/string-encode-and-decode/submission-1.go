type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := make([]byte, 0)

	for _, str := range strs {
		res = append(res, str...)
		res = append(res, 0)
	}

	return string(res)
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	temp := make([]byte, 0)

	for _, c := range encoded {
		if byte(c) == 0 {
			res = append(res, string(temp))
			temp = make([]byte, 0)
			continue
		}

		temp = append(temp, byte(c))
	}

	return res
}
