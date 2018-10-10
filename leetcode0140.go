func wordBreak(s string, wordDict []string) []string {
    lengths := make(map[int]bool)
    fails := make(map[int]bool)
    words := make(map[string]bool)
    for _, w := range wordDict {
        lengths[len(w)] = true
        words[w] = true
    }
    results := []string {}
    dfs(s, 0, []int {}, lengths, fails, words, &results)
    return results
}

func dfs(s string, start int, starts []int, lengths, fails map[int]bool, words map[string]bool, results *[]string) bool {
    if fails[start] {
        return false
    }
    fail := true
    starts = append(starts, start)
    if words[s[start:]] {
        result := ""
        for i := 0; i < len(starts) - 1; i++ {
            result += s[starts[i]:starts[i + 1]] + " "
        }
        result += s[start:]
        *results = append(*results, result)
        fail = false
    }
    for l, _ := range lengths {
        if start + l > len(s) || !words[s[start:start + l]] {
            continue
        }
        if dfs(s, start + l, starts, lengths, fails, words, results) {
            fail = false
        }
    }
    starts = starts[:len(starts) - 1]
    if fail {
        fails[start] = true
    }
    return !fail
}