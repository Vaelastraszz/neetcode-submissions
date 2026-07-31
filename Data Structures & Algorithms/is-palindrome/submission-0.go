func isPalindrome(s string) bool {
	
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned_str := strings.ToLower(re.ReplaceAllString(s, ""))
	
	for left, right := 0, len(cleaned_str) - 1; left <= right; {
		if cleaned_str[left] != cleaned_str[right] {
			return false
		}
		left ++
		right --
	}
	return true
}
