func longestCommonPrefix(strs []string) string {
    
	longestPref := "" 
	firstW := strs[0]
	nWords := len(strs)

	mainLoop: for i := range firstW {
		pref := firstW[0:i+1]

		for j:=1; j<nWords; j++ {
			word := strs[j]
			if len(word) < len(pref) || pref != word[0:i+1] {
				break mainLoop
			}
		}

		longestPref = pref

	}
	
	return longestPref
}

