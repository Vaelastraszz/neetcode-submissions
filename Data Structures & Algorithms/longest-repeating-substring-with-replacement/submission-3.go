type Occurences map[rune]int

func newOccurences() Occurences {
	return make(map[rune]int)
}

func (o Occurences) MaxOcc() rune {
	topFreq := 0
	var topLetter rune
	
	for k, v := range o {

		if o[k] > topFreq {
			topLetter = k
			topFreq = v
		}

	}

	return topLetter
}

func characterReplacement(s string, k int) int {

	left, size, length := 0, 0, 0
	newS := []rune(s)

	topLetter := newS[left]
	mapOccurences := newOccurences()

	for right, letter := range newS {

		size = right - left + 1 
		mapOccurences[letter]++

		if letter != topLetter && mapOccurences[letter] > mapOccurences[topLetter] {

			topLetter = letter

		}

		if size - mapOccurences[topLetter] <= k {

			length = max(length, size)

		} else {

			mapOccurences[newS[left]]--
			left ++
			topLetter = mapOccurences.MaxOcc()
			
		}

	}

	return length
}

