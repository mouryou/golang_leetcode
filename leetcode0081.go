func search(nums []int, target int) bool {
    ln := len(nums)
    l, r := 0, ln
    for l < r {
        m := (l + r) / 2
        if nums[m] == target || nums[l] == target {
            return true
        }
        if m == l {
            return false
        }
        if nums[m] > nums[l] {
            if target < nums[m] && target > nums[l] {
                r = m
            } else {
                l = m + 1
            }
        } else if nums[m] < nums[l] {
            if target < nums[l] && target > nums[m] {
                l = m + 1
            } else {
                r = m
            }
        } else if nums[r - 1] < nums[l] {
            l = m + 1
        } else {
            return search(nums[l + 1:m], target) || search(nums[m + 1:r], target)
        }
    }
    return false
}