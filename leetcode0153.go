func findMin(nums []int) int {
    l, r := 0, len(nums)
    for l < r {
        if nums[l] < nums[r - 1] {
            return nums[l]
        }
        m := (l + r) / 2
        if m == 0 || nums[m] < nums[m - 1] {
            return nums[m]
        }
        if nums[m] > nums[l] {
            l = m + 1
        } else {
            r = m
        }
        
    }
    return nums[l]
}