func minCut(s string) int {
    l := len(s)
    mins := make([]int, l)
    dp := make([][]bool, l)
    for i := 0; i < l; i++ {
        dp[i] = make([]bool, i + 1)
    }
    for i := 0; i < l; i++ {
        min := i
        for j := 0; j <= i; j++ {
            if s[i] == s[j] && (j + 1 >= i - 1 || dp[i - 1][j + 1]) {
                dp[i][j] = true
                if j == 0 {
                    min = 0
                    break
                } else {
                    if mins[j - 1] + 1 < min {
                        min = mins[j - 1] + 1
                    }
                }
            }
        }
        mins[i] = min
    }
    return mins[l - 1]
}
