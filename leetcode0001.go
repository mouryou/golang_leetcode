func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, n := range nums {
        if j, ok := m[target - n]; ok {
            return []int{j, i}
        } else {
            m[n] = i
        }
    }
    return []int{-1, -1}
}