func jump(nums []int) int {
    l := len(nums)
    if l == 1 {
        return 0
    }
    dp := make([]int, l)
    for i := 1; i < l; i++ {
        dp[i] = -1
    }
    for i, d := range nums { // index, distance
        r := i + d
        if r >= l - 1 {
            return dp[i] + 1
        }
        for j := r; j > i && dp[j] == -1; j-- {
            dp[j] = dp[i] + 1
        }
    }
    return dp[l - 1]
}