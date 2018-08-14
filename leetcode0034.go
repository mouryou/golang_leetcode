func searchRange(nums []int, target int) []int {
    l, r := 0, len(nums)
    res := []int {-1, -1}
    for l < r {
        m := (l + r) / 2
        if nums[m] > target {
            r = m
        } else if nums[m] < target {
            l = m + 1
        } else if m == 0 || nums[m - 1] < target {
            res[0] = m
            break
        } else {
            r = m
        }
    }
    if res[0] == -1 {
        return res
    }
    l, r = res[0], len(nums)
    for l < r {
        m := (l + r) / 2
        if nums[m] > target {
            r = m
        } else if nums[m] < target {
            l = m + 1
        } else if m == len(nums) - 1 || nums[m + 1] > target {
            res[1] = m
            return res
        } else {
            l = m + 1
        }
    }
    return res
}