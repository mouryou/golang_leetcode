func generateParenthesis(n int) []string {
    dp := make([]map[string]bool, n + 1)
    dp[0] = map[string]bool {"":true}
    for i := 1; i <= n; i++ {
        dp[i] = make(map[string]bool)
        for s, _ := range dp[i - 1] {
            dp[i]["(" + s + ")"] = true
        }
        for j := 1; j < i; j++ {
            for s1, _ := range dp[j] {
                for s2, _ := range dp[i - j] {
                    dp[i][s1 + s2] = true
                }
            }
        }
    }
    keys := make([]string, len(dp[n]))
    i := 0
    for k, _ := range dp[n] {
        keys[i] = k
        i += 1
    }
    return keys
}
