func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    first := strs[0]

    for i := range first {
        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || first[i] != strs[j][i] {
                return first[:i]
            }
        }
    }

    return first
}