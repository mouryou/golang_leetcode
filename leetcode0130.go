func solve(board [][]byte)  {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        dfs(i, 0, board)
        dfs(i, n - 1, board)
    }
    for j := 1; j < len(board[0]) - 1; j++ {
        dfs(0, j, board)
        dfs(m - 1, j, board)
    }
    for i, row := range board {
        for j, cell := range row {
            if cell == 'I' {
                board[i][j] = 'O'
            } else if cell == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}

func dfs(i, j int, board [][]byte) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
        return
    }
    board[i][j] = 'I'
    directions := [][]int {[]int {1, 0}, []int {-1, 0}, []int {0, 1}, []int {0, -1}}
    for _, d := range directions {
        ni, nj := i + d[0], j + d[1]
        dfs(ni, nj, board)
    }
}