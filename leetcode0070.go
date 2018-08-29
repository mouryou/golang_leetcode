func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    dp1, dp2 := 1, 2
    for i := 3; i <= n; i++ {
        dp1, dp2 = dp2, dp1 + dp2
    }
    return dp2
}