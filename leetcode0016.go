func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    res := nums[0] + nums[1] + nums[2]
    i := 0
    l := len(nums)
    for i < l - 2 {
        s := nums[i] + nums[i + 1] + nums[i + 2]
        if s > target && s - target >= abs(res - target) {
            return res
        }
        
        s = nums[i] + nums[l - 2] + nums[l - 1]
        if !(s < target && target - s >= abs(res - target)) {
            j, k := i + 1, l - 1
            for j < k {
                s = nums[i] + nums[j] + nums[k]
                if s > target {
                    if s - target < abs(res - target) {
                        res = s
                    }
                    k -= 1
                    for k > j && nums[k] == nums[k + 1] {
                        k -= 1
                    }
                } else if s < target {
                    if target - s < abs(res - target) {
                        res = s
                    }
                    j += 1
                    for j < k && nums[j] == nums[j - 1] {
                        j += 1
                    }
                } else {
                    return target
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