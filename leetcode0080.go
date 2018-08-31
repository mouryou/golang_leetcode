func removeDuplicates(nums []int) int {
    l := len(nums)
    if l < 3 {
        return l
    }
    i := 1
    for j := 2; j < l; j++ {
        if nums[j] != nums[i] || nums[j] != nums[i - 1] {
            i++
            if i != j {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
    return i + 1
}