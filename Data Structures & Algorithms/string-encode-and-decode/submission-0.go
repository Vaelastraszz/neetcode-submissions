type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := []byte{}

	for _, str := range strs {
		result = append(result, strconv.Itoa(len(str))...)
		result = append(result, '!')
		result = append(result, str...)
	}

	return string(result)
}

func (s *Solution) Decode(encoded string) []string {

	result := []string{}
	i := 0

	for i < len(encoded) {

		j := i

		for encoded[j] != '!' {
			j ++ 
		}

		length, _ := strconv.Atoi(encoded[i:j])
		result  = append(result, encoded[j+1: j+1+length])

		i = j+1+length
	}
	
	return result
}
