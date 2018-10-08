func singleNumber(nums []int) int {
    a, b := 0, 0
    // 0 0 0 -> 0 0
    // 0 0 1 -> 0 1
    // 0 1 0 -> 0 1
    // 0 1 1 -> 1 0
    // 1 0 0 -> 1 0
    // 1 0 1 -> 0 0
    for _, n := range nums {
        a, b = (a ^ b) & (a ^ n), (b ^ n) & ^a
    }
    return b
}