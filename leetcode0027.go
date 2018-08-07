func removeElement(nums []int, val int) int {
    i := 0
    for j := 0; j < len(nums); j++ {
        if nums[j] != val {
            if j != i {
                nums[i] = nums[j]
            }
            i++
        }
    }
    return i
}