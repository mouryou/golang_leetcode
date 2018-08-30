func sortColors(nums []int)  {
    l := len(nums)
    if l < 2 {
        return
    }
    b0, b2 := 0, l - 1
    for b0 < l && nums[b0] == 0 {
        b0++
    }
    if b0 == l {
        return
    }
    for b2 >= 0 && nums[b2] == 2 {
        b2--
    }
    if b2 == -1 {
        return
    }
    i := b0
    for i <= b2 {
        if nums[i] == 0 {
            nums[i], nums[b0] = nums[b0], nums[i]
            for nums[b0] == 0 {
                b0++
            }
            if i < b0 {
                i = b0
            }
        } else if nums[i] == 2 {
            nums[i], nums[b2] = nums[b2], nums[i]
            for nums[b2] == 2 {
                b2--
            }
        } else {
            i++
        }
    }
}