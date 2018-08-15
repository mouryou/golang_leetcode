func isValidSudoku(board [][]byte) bool {
    // examine every 3x3 grid
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            numbers := make(map[byte]bool)
            for ii := 0; ii < 3; ii++ {
                for jj := 0; jj < 3; jj++ {
                    s := board[i * 3 + ii][j * 3 + jj]
                    if s == '.' {
                        continue
                    }
                    if _, ok := numbers[s]; ok {
                        return false
                    } else {
                        numbers[s] = true
                    }
                    
                }
            }
        }
    }
    // examine every row
    for i := 0; i < 9; i++ {
        numbers := make(map[byte]bool)
        for j := 0; j < 9; j++ {
            s := board[i][j]
            if s == '.' {
                continue
            }
            if _, ok := numbers[s]; ok {
                return false
            } else {
                numbers[s] = true
            }
        }
    }
    // examine every column
    for j := 0; j < 9; j++ {
        numbers := make(map[byte]bool)
        for i := 0; i < 9; i++ {
            s := board[i][j]
            if s == '.' {
                continue
            }
            if _, ok := numbers[s]; ok {
                return false
            } else {
                numbers[s] = true
            }
        }
    }
    return true
}