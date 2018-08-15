func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        m := (l + r) / 2
        if nums[m] > target {
            r = m
        } else if nums[m] < target {
            l = m + 1
        } else {
            return m
        }
    }
    if l == len(nums) {
        return l
    }
    if nums[l] < target {
        return l + 1
    } else {
        return l
    }
}