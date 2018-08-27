func permuteUnique(nums []int) [][]int {
    res := make([][]int, 0)
    permute(nums, 0, &res)
    return res
}

func permute(nums []int, index int, result *[][]int) {
    if index == len(nums) {
        *result = append(*result, append([]int {}, nums...))
        return
    }
    appeared := make(map[int]bool)
    for i := index; i < len(nums); i++ {
        if appeared[nums[i]] {
            continue
        }
        nums[index], nums[i] = nums[i], nums[index]
        permute(nums, index + 1, result)
        appeared[nums[index]] = true
        nums[index], nums[i] = nums[i], nums[index]
    }
    
}