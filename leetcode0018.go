func fourSum(nums []int, target int) [][]int {
    res := make([][]int, 0)
    l := len(nums)
    if l < 4 {
        return res
    }
    sort.Ints(nums)
    for i := 0; i < l - 3; i += 1 {
        if (i > 0 && nums[i] == nums[i - 1]) || nums[i] + nums[l - 3] + nums[l - 2] + nums[l - 1] < target {
            continue
        }
        if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target {
            return res
        }
        for j := i + 1; j < l - 2; j += 1 {
            if (j > i + 1 && nums[j] == nums[j - 1]) || nums[i] + nums[j] + nums[l - 2] + nums[l - 1] < target {
                continue
            }
            if nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target {
                break
            }
            m, n := j + 1, l - 1
            for m < n {
                s := nums[i] + nums[j] + nums[m] + nums[n]
                if s > target {
                    n -= 1
                    for n > m && nums[n] == nums[n + 1] {
                        n -= 1
                    }
                } else if s < target {
                    m += 1
                    for m < n && nums[m] == nums[m - 1] {
                        m += 1
                    }
                } else {
                    res = append(res, []int {nums[i], nums[j], nums[m], nums[n]})
                    n -= 1
                    for n > m && nums[n] == nums[n + 1] {
                        n -= 1
                    }
                    m += 1
                    for m < n && nums[m] == nums[m - 1] {
                        m += 1
                    }
                }
            }
        }
    } 
    return res
}