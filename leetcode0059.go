func generateMatrix(n int) [][]int {
    if n == 0 {
        return make([][]int, 0)
    }
    result := make([][]int, n)
    for i := 0; i < n; i++ {
        result[i] = make([]int, n)
    }
    deltas := [][]int {{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    direction := 0
    i, j := 0, 0
    for k := 1; k <= n * n; k++ {
        result[i][j] = k
        nextI, nextJ := i + deltas[direction][0], j + deltas[direction][1]
        if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= n || result[nextI][nextJ] != 0 {
            direction = (direction + 1) % 4
        }
        i, j = i + deltas[direction][0], j + deltas[direction][1]
    }
    return result
}