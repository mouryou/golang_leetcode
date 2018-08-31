func exist(board [][]byte, word string) bool {
    m := len(board)
    if m == 0 {
        return false
    }
    n := len(board[0])
    if n == 0 {
        return false
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == word[0] {
                board[i][j] = 0
                if dfs(board, word, i, j, 1) {
                    return true
                }
                board[i][j] = word[0]
            }
        }
    }
    return false
}

func dfs(board [][]byte, word string, i, j, wi int) bool {
    if wi == len(word) {
        return true
    }
    for _, delta := range [4][2]int {{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
        nexti, nextj := i + delta[0], j + delta[1]
        if nexti >= 0 && nexti < len(board) && nextj >= 0 && nextj < len(board[0]) && word[wi] == board[nexti][nextj] {
            board[nexti][nextj] = 0
            if dfs(board, word, nexti, nextj, wi + 1) {
                return true
            }
            board[nexti][nextj] = word[wi]
        }
    }
    return false
}