func firstMissingPositive(nums []int) int {
    pe := 0 // one positive element
    for _, n := range nums {
        if n > 0 {
            pe = n
            break
        }
    }
    if pe == 0 {
        return 1
    }
    for i, n := range nums {
        if n <= 0 {
            nums[i] = pe
        }
    }
    for _, n := range nums {
        if n < 0 {
            n = -n
        }
        if n <= len(nums) && nums[n - 1] > 0 {
            nums[n - 1] = -nums[n - 1]
        }
    }
    fmt.Println(nums)
    for i := 1; i <= len(nums); i++ {
        if nums[i - 1] > 0 {
            return i
        }
    }
    return len(nums) + 1
}
