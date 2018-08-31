func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    if m == 0 {
        return 0
    }
    n := len(matrix[0])
    if n == 0 {
        return 0
    }
    heights := make([]int, n)
    lefts := make([]int, n)
    rights := make([]int, n)
    res := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                heights[j] += 1
            } else {
                heights[j] = 0
            }
            lefts[j] = 0
            rights[j] = n
        }
        stack := []int {}
        for j := n - 1; j >= 0; j-- {
            ls := len(stack)
            for ls > 0 && heights[stack[ls - 1]] > heights[j] {
                lefts[stack[ls - 1]] = j + 1
                stack = stack[:ls - 1]
                ls--
            }
            stack = append(stack, j)
        }
        stack = []int {}
        for j := 0; j < n; j++ {
            ls := len(stack)
            for ls > 0 && heights[stack[ls - 1]] > heights[j] {
                rights[stack[ls - 1]] = j
                stack = stack[:ls - 1]
                ls--
            }
            stack = append(stack, j)
        }
        for j := 0; j < n; j++ {
            area := (rights[j] - lefts[j]) * heights[j] 
            if area > res {
                res = area
            }
        }
    }
    return res
}