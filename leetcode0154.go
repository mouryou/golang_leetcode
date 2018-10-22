func findMin(nums []int) int {
    l, r := 0, len(nums)
    for l < r {
        if r - l == 1 {
            return nums[l]
        }
        m := (l + r) / 2
        if m == 0 || nums[m] < nums[m - 1] {
            return nums[m]
        }
        if nums[l] < nums[r - 1] {
            return nums[l]
        } else if nums[l] > nums[r - 1] {
            if nums[m] >= nums[l] {
                l = m + 1
            } else {
                r = m
            }
        } else {
            if nums[m] > nums[l] {
                l = m + 1
            } else if nums[m] < nums[l] {
                r = m
            } else {
                a, b := findMin(nums[l:m]), findMin(nums[m:r])
                if a > b {
                    return b
                } else {
                    return a
                }
            }
        }
    }
    return nums[l]
}