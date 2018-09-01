func subsetsWithDup(nums []int) [][]int {
    count := make(map[int]int)
    for _, n := range nums {
        count[n] += 1
    }
    keys := make([]int, 0)
    for k, _ := range count {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    res := [][]int {[]int {}}
    for _, k := range keys {
        l := len(res)
        for i := 0; i < l; i++ {
            r := make([]int, len(res[i]))
            copy(r, res[i])
            for j := 0; j < count[k]; j++ {
                r = append(append([]int {}, r...), k)
                res = append(res, r)
            }
        }
    }
    return res
}