func solveSudoku(board [][]byte)  {
    solveSudokuRecursive(board, 0)
}

// pos == i * 9 + j, range: 0-80
func solveSudokuRecursive(board [][]byte, pos int) bool {
    i, j := pos / 9, pos % 9
    if board[i][j] == '.' {
        choices := make(map[byte]bool)
        for k := 0; k < 9; k++ {
            choices[board[i][k]] = true
            choices[board[k][j]] = true
            choices[board[i / 3 * 3 + k / 3][j / 3 * 3 + k % 3]] = true
        }
        for _, k := range "123456789" {
            if !choices[byte(k)] {
                board[i][j] = byte(k)
                if pos == 80 || solveSudokuRecursive(board, pos + 1) {
                    return true
                }
            }
        }
        board[i][j] = '.'
        return false
    } else {
        return pos == 80 || solveSudokuRecursive(board, pos + 1)
    }
} 