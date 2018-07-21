func isMatch(s string, p string) bool {
    ls := len(s)
    lp := len(p)
    dp := make([]bool, lp + 1)
    dp_next := make([]bool, lp + 1)
    dp[0] = true
    var i, j int
    for j = 1; j <= lp; j++ {
        if p[j - 1] == '*' {
            dp[j] = dp[j - 2]
        } else {
            dp[j] = false
        }
    }
    for i = 1; i <= ls; i++ {
        dp_next[0] = false
        for j = 1; j <= lp; j++ {
            if p[j - 1] == '.' {
                dp_next[j] = dp[j - 1]
            } else if p[j - 1] == '*' {
                dp_next[j] = dp_next[j - 2] || ((s[i - 1] == p[j - 2] || p[j - 2] == '.') && dp[j])
            } else if p[j - 1] == s[i - 1] {
                dp_next[j] = dp[j - 1]
            } else {
                dp_next[j] = false
            }
        }
        dp, dp_next = dp_next, dp
    }
    return dp[lp]
}