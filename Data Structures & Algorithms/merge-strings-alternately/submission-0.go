func mergeAlternately(word1 string, word2 string) string {
	res := []rune{}
	w1 := []rune(word1)
	w2 := []rune(word2)

	end1, end2 := len(w1) -1, len(w2) - 1

	// on cherche calcule la fin la plus longue
	maxEnd := max(end1, end2)
	// on boucle de 0 a la fin la plus longue
	for current := 0; current <= maxEnd; current ++ {
	// Si le pointer est en dessous de end1 pour w1 on ajoute la lettre
		if current <= end1 {
			res = append(res, w1[current])
		}
	// Meme chose pour w2
		if current <= end2 {
			res = append(res, w2[current])
		}
	}
	// on retourne a la fin le mot comple en string
	return string(res)
}
