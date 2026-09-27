func longestPalindrome(s string) string {
    runes := []rune(s)
    var subOdd, subEven string

    for i := range runes {
        tmpSubOdd := getSubPal(runes, i, i)
        if len([]rune(tmpSubOdd)) > len([]rune(subOdd)) {
            subOdd = tmpSubOdd
        }

        tmpSubEven := getSubPal(runes, i, i+1)
        if len([]rune(tmpSubEven)) > len([]rune(subEven)) {
            subEven = tmpSubEven
        }
    }

    if len([]rune(subOdd)) > len([]rune(subEven)) {
        return subOdd
    }

    return subEven
}

func getSubPal(haystack []rune, l, r int) string {
    for l >= 0 && r < len(haystack) && haystack[l] == haystack[r] {
        l--
        r++
    }

    return string(haystack[l+1 : r])
}