func combine(n int, k int) [][]int {
    if k > n {
        return make([][]int, 0)
    }
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    return combineSlice(nums, k)
}

func combineSlice(nums []int, k int) [][]int {
    if k == 0 {
        return [][]int {[]int {}}
    }
    l := len(nums)
    if l == k {
        return [][]int {nums}
    }
    res := make([][]int, 0)
    for i := 0; i <= l - k; i++ {
        for _, results := range combineSlice(nums[i + 1:], k - 1) {
            res = append(res, append([]int {nums[i]}, results...))
        }
    }
    return res
}