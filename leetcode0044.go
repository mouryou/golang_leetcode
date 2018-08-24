func isMatch(s string, p string) bool {
    ls, lp := len(s), len(p)
    dp := make([]bool, ls + 1)
    dp[0] = true
    dp_next := make([]bool, ls + 1)
    for i := 1; i <= lp; i++ {
        if p[i - 1] == '*' {
            if i > 1 && p[i - 2] == '*' {
                continue
            }
            dp_next[0] = dp[0]
            for j := 1; j <= ls; j++ {
                dp_next[j] = dp_next[j - 1] || dp[j]
            }
        } else {
            dp_next[0] = false
            for j := 1; j <= ls; j++ {
                dp_next[j] = dp[j - 1] && (p[i - 1] == s[j - 1] || p[i - 1] == '?')
            }
        }
        dp, dp_next = dp_next, dp
    }
    return dp[ls]
}