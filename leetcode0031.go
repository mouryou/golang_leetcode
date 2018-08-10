func nextPermutation(nums []int)  {
    i := len(nums) - 1
    for i > 0 && nums[i] <= nums[i - 1] {
        i--
    }
    if i > 0 {
        j := i
        for j < len(nums) - 1 && nums[i - 1] < nums[j + 1] {
            j++
        }
        nums[i - 1], nums[j] = nums[j], nums[i - 1]
    }
    j := len(nums) - 1
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
