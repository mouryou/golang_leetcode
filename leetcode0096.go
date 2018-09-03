func numTrees(n int) int {
    dp := make(map[[2]int]int)
    for length := 1; length <= n; length++ {
        for l := 1; l <= n - length + 1; l++ {
            r := l + length
            res := 0
            for mid := l; mid < r; mid++ {
                var nl, nr int
                if mid == l {
                    nl = 1
                } else {
                    nl = dp[[2]int {l, mid}]
                }
                if mid == r - 1 {
                    nr = 1
                } else {
                    nr = dp[[2]int {mid + 1, r}]
                }
                res += nl * nr
            }
            dp[[2]int {l, r}] = res
        }
    }
    return dp[[2]int {1, n + 1}]
}