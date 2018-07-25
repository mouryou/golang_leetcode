func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    for i, r := range strs[0] {
        for j := 1; j < len(strs); j++ {
            if len(strs[j]) <= i || strs[j][i] != byte(r) {
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}