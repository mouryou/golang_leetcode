func numDistinct(s string, t string) int {
    ls, lt := len(s), len(t)
    dp := make([]int, ls + 1)
    next_dp := make([]int, ls + 1) 
    for i, _ := range dp {
        dp[i] = 1
    }
    for i := 1; i <= lt; i++ {
        next_dp[0] = 0
        for j := 1; j <= ls; j++ {
            next_dp[j] = next_dp[j - 1]
            if t[i - 1] == s[j - 1] {
                next_dp[j] += dp[j - 1]
            }
        }
        dp, next_dp = next_dp, dp
    }
    return dp[ls]
}