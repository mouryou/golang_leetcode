func canJump(nums []int) bool {
    l := len(nums)
    dp := make([]bool, l)
    dp[0] = true
    for i := 0; i < l; i++ {
        if !dp[i] {
            return false
        }
        if i + nums[i] >= l - 1 {
            return true
        }
        for j := i + nums[i]; j > i; j-- {
            if !dp[j] {
                dp[j] = true
            } else {
                break
            }
        }
    }
    return dp[l - 1]
}