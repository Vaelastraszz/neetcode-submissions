func countSubstrings(s string) int {
    runes := []rune(s)
    var count int

    var palCount func(haystack []rune, l, r int)

    palCount = func(haystack []rune, l, r int) {

        for l >= 0 && r < len(haystack) && haystack[l] == haystack[r] {
        count ++ 
        l--
        r++
    }

    }

    for i := range runes {

        palCount(runes, i, i)
        palCount(runes, i, i+1)
  
    }

    return count
}

