func wordBreak(s string, wordDict []string) bool {
    lengths := make(map[int]bool)
    results := make(map[int]bool)
    words := make(map[string]bool)
    for _, w := range wordDict {
        lengths[len(w)] = true
        words[w] = true
    }
    return dfs(s, 0, lengths, words, results)
}

func dfs(s string, start int, lengths map[int]bool, words map[string]bool, results map[int]bool) bool {
    if words[s[start:]] {
        return true
    }
    if _, ok := results[start]; ok {
        return false
    }
    for l, _ := range lengths {
        if l > len(s) - start || !words[s[start:start + l]] {
            continue
        }
        if dfs(s, start + l, lengths, words, results) {
            return true
        }
    }
    results[start] = false
    return false
}