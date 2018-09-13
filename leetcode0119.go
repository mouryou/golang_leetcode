func getRow(rowIndex int) []int {
    res := make([]int, rowIndex + 1)
    res[0] = 1
    for i := 1; i <= rowIndex / 2; i++ {
        res[i] = 1
        for j := i + 1; j <= rowIndex; j++ {
            res[i] *= j
            res[i] /= j - i
        }
    }
    for i := rowIndex / 2 + 1; i <= rowIndex; i++ {
        res[i] = res[rowIndex - i]
    }
    return res
}