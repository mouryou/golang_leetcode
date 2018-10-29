func findPeakElement(nums []int) int {
    l, r := 0, len(nums)
    for l < r {
        if isPeak(nums, l) {
            return l
        }
        if isPeak(nums, r - 1) {
            return r - 1
        }
        m := (l + r) / 2
        if isPeak(nums, m) {
            return m
        }
        if (nums[m] > nums[l] && nums[m] > nums[r - 1]) || (nums[m] < nums[l] && nums[m] < nums[r - 1]) {
            if nums[m] <= nums[m - 1] {
                r = m
            } else {
                l = m + 1
            }
        } else if nums[m] > nums[l] {
            l = m + 1
        } else {
            r = m
        }
    }
    return l
}

func isPeak(nums []int, i int) bool {
    return (i == 0 || nums[i] > nums[i - 1]) && (i == len(nums) - 1 || nums[i] > nums[i + 1])
}