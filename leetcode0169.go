func majorityElement(nums []int) int {
    curNum, curCnt := nums[0], 1
    for i := 1; i < len(nums); i++ {
        if nums[i] != curNum {
            if curCnt > 0 {
                curCnt--
            } else {
                curNum = nums[i]
                curCnt = 1
            }
        } else {
            curCnt++
        }
    }
    return curNum
}