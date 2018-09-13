func minimumTotal(triangle [][]int) int {
    l := len(triangle)
    if l == 0 {
        return 0
    }
    dp := make([]int, l)
    dp[0] = triangle[0][0]
    next_dp := make([]int, l)
    for i := 1; i < l; i++ {
        next_dp[0] = triangle[i][0] + dp[0]
        for j := 1; j <= i - 1; j++ {
            next_dp[j] = dp[j - 1]
            if dp[j] < next_dp[j] {
                next_dp[j] = dp[j]
            }
            next_dp[j] += triangle[i][j]
        }
        next_dp[i] = triangle[i][i] + dp[i - 1]
        dp, next_dp = next_dp, dp
    }
    res := dp[0]
    for i := 1; i < l; i++ {
        if dp[i] < res {
            res = dp[i]
        }
    }
    return res
}