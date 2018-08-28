func maxSubArray(nums []int) int {
    minSum := 0
    result := 0
    curSum := 0
    max := nums[0]
    for _, n := range nums {
        curSum += n
        if curSum - minSum > result {
            result = curSum - minSum
        }
        if curSum < minSum {
            minSum = curSum
        }
        if n > max {
            max = n
        }
    }
    if max < 0 {
        return max
    }
    return result
}