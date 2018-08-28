func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    if m == 0 {
        return []int {}
    }
    n := len(matrix[0])
    if n == 0 {
        return []int {}
    }
    result := make([]int, m * n)
    delta := [4][2]int {[2]int {0, 1}, [2]int {1, 0},[2]int {0, -1},[2]int {-1, 0}}
    direction := 0
    i, j := 0, 0
    printed := make(map[[2]int]bool)
    for t := 0; t < m * n; t++ {
        result[t] = matrix[i][j]
        printed[[2]int{i, j}] = true
        nextI, nextJ := i + delta[direction][0], j + delta[direction][1]
        if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n || printed[[2]int{nextI, nextJ}] {
            direction = (direction + 1) % 4
            nextI, nextJ = i + delta[direction][0], j + delta[direction][1]
        }
        i, j = nextI, nextJ
    }
    return result
}