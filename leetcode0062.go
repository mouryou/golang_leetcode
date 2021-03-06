func uniquePaths(m int, n int) int {
    dp := make([]int, m)
    for i := 0; i < m; i++ {
        dp[i] = 1
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[j] += dp[j - 1]
        }
    }
    return dp[m - 1]
}