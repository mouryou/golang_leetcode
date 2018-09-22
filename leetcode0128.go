func longestConsecutive(nums []int) int {
    ls := make(map[int]int)
    res := 0
    for _, n := range nums {
        if ls[n] == 0 {
            ll, lr := ls[n - 1], ls[n + 1]
            ln := ll + lr + 1
            if ln > res {
                res = ln
            }
            ls[n] = ln
            ls[n - ll] = ln
            ls[n + lr] = ln
        }
    }
    return res
}