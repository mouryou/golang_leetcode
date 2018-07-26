func threeSum(nums []int) [][]int {
    res := make([][]int, 0)
    l := len(nums)
    sort.Ints(nums)
    i := 0
    for i < l - 2 {
        if nums[i] + nums[i + 1] + nums[i + 2] > 0 {
            return res
        }
        if nums[i] + nums[l - 2] + nums[l - 1] >= 0 {
            j, k := i + 1, l - 1
            for j < k {
                s := nums[i] + nums[j] + nums[k]
                if s < 0 {
                    j += 1
                    for j < k && nums[j] == nums[j - 1] {
                        j += 1
                    }
                } else if s > 0 {
                    k -= 1
                    for k > j && nums[k] == nums[k + 1] {
                        k -= 1
                    }
                } else {
                    res = append(res, []int{nums[i], nums[j], nums[k]})
                    j += 1
                    for j < k && nums[j] == nums[j - 1] {
                        j += 1
                    }
                    k -= 1
                    for k > j && nums[k] == nums[k + 1] {
                        k -= 1
                    }
                }
            }
        }
        i += 1
        for i < l - 2 && nums[i] == nums[i - 1] {
            i += 1
        }
        
    }
    return res
}