func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    if obstacleGrid[0][0] == 1 || obstacleGrid[m - 1][n - 1] == 1 {
        return 0
    }
    obstacleGrid[0][0] = -1
    for i := 1; i < n; i++ {
        if obstacleGrid[0][i] == 1 {
            break
        } else {
            obstacleGrid[0][i] = -1
        }
    }
    for i := 1; i < m; i++ {
        if obstacleGrid[i][0] != 1 && obstacleGrid[i - 1][0] != 1 {
            obstacleGrid[i][0] = obstacleGrid[i - 1][0]
        }
        for j := 1; j < n; j++ {
            if obstacleGrid[i][j] != 1 {
                if obstacleGrid[i - 1][j] != 1 {
                    obstacleGrid[i][j] += obstacleGrid[i - 1][j]
                }
                if obstacleGrid[i][j - 1] != 1 {
                    obstacleGrid[i][j] += obstacleGrid[i][j - 1]
                }
            }
        }
    }
    return -obstacleGrid[m - 1][n - 1]
}