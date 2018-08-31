func subsets(nums []int) [][]int {
    if len(nums) == 0 {
        return [][]int {[]int {}}
    }
    res := [][]int {[]int {}, []int {nums[0]}}
    for _, subset := range subsets(nums[1:]) {
        if len(subset) > 0 {
            res = append(res, subset)
            res = append(res, append([]int {nums[0]}, subset...))
        }
    }
    return res
}