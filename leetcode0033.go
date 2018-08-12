func search(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        m := (l + r) / 2
        if nums[m] == target {
            return m
        }
        if ((nums[m] > nums[0]) != (target < nums[m])) != (target >= nums[0]) {
            r = m
        } else {
            l = m + 1
        }
    }
    return -1
}