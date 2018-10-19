func maxProduct(nums []int) int {
    l := len(nums)
    res, max, min := nums[0], nums[0], nums[0]
    for i := 1; i < l; i++ {
        n := nums[i]
        max, min = max2(max2(n, max * n), min *n), min2(min2(n, max * n), min *n)
        if max > res {
            res = max
        }
    }
    return res
}

func max2(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min2(a, b int) int {
    if a > b {
        return b
    }
    return a
}