func grayCode(n int) []int {
    res := []int {0}
    for i := 0; i < n; i++ {
        l := len(res)
        for j := l - 1; j >= 0; j-- {
            res = append(res, l + res[j])
        }
    }
    return res
}