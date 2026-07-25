func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false 
	}
	r1 := []rune(s)
	r2 := []rune(t)

	size_str := len(r1)
	mapWords := make(map[string]map[rune]int, 2)
	map_w1 := make(map[rune]int, size_str)
	map_w2 := make(map[rune]int, size_str)

	mapWords[s] = map_w1
	mapWords[t] = map_w2

	for i:= 0; i < size_str; i ++ {
			mapWords[s][r1[i]] ++ 
			mapWords[t][r2[i]] ++ 	
	}

	for letter, occ_map_1 := range mapWords[s] {
		if occ_map_2, ok := mapWords[t][letter]; ok {
			if occ_map_1 != occ_map_2 {
				return false
			}
			continue
		}
		return false
	}

	return true 
}
