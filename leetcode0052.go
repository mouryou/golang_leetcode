func totalNQueens(n int) int {
    // conflict:
    // 0 ~ n-1: column j
    // n ~ 3n-2: diagonal 2n-1+j-i
    // 3n-1 ~ 5n-3: anti-diagonal 3n-1+i+j
    result := 0
    search(n, 0, make([]bool, 5 * n - 2), &result)
    return result
}

func search(n, i int, conflict []bool, result *int) {
    if i == n {
        *result += 1
        return
    }
    for j := 0; j < n; j++ {
        if conflict[j] || conflict[2 * n + j - i] || conflict[3 * n - 1 + i + j] {
            continue
        }
        conflict[j] = true
        conflict[2 * n + j - i] = true
        conflict[3 * n - 1 + i + j] = true
        search(n, i + 1, conflict, result)
        conflict[j] = false
        conflict[2 * n + j - i] = false
        conflict[3 * n - 1 + i + j] = false
    }
}