func solveNQueens(n int) [][]string {
    // 0 ~ n-1: column j
    // n ~ 3n-2: diagonal 2n-1+j-i
    // 3n-1 ~ 5n-3: anti-diagonal 3n-1+i+j
    rows := make([]string, n)
    for i := 0; i < n; i++ {
        rows[i] = strings.Repeat(".", i) + "Q" + strings.Repeat(".", n - 1 - i)
    }
    conflict := make([]bool, 5 * n - 2)
    cur_board := make([]string, n)
    results := make([][]string, 0)
    search(n, 0, cur_board, conflict, &results, rows)
    return results
    
}

func search(n, i int, cur_board []string, conflict []bool, results *[][]string, rows []string) {
    if i == n {
        *results = append(*results, append([]string {}, cur_board...))
        return
    }
    for j := 0; j < n; j++ {
        if conflict[j] || conflict[2 * n - 1 + j - i] || conflict[3 * n - 1 + i + j] {
            continue
        }
        conflict[j] = true
        conflict[2 * n - 1 + j - i] = true
        conflict[3 * n - 1 + i + j] = true
        cur_board[i] = rows[j]
        search(n, i + 1, cur_board, conflict, results, rows)
        conflict[j] = false
        conflict[2 * n - 1 + j - i] = false
        conflict[3 * n - 1 + i + j] = false
    } 
}