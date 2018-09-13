func generate(numRows int) [][]int {
    if numRows == 0 {
        return [][]int {}
    }
    res := [][]int {[]int {1}}
    for i := 1; i < numRows; i++ {
        l := make([]int, i + 1)
        copy(l, res[i - 1])
        for j, n := range res[i - 1] {
            l[j + 1] += n
        }
        res = append(res, l)
    }
    return res
}