func setZeroes(matrix [][]int)  {
    m := len(matrix)
    if m == 0 {
        return
    }
    n := len(matrix[0])
    if n == 0 {
        return
    }
    c0, r0 := false, false
    if matrix[0][0] == 0 {
        c0, r0 = true, true
    } else {
        for i := 1; i < m; i++ {
            if matrix[i][0] == 0 {
                c0 = true
                break
            }
        }
        for j := 1; j < n; j++ {
            if matrix[0][j] == 0 {
                r0 = true
                break
            }
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if c0 {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
    if r0 {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    
}