func largestRectangleArea(heights []int) int {
    lh := len(heights)
    left, right := make([]int, lh), make([]int, lh)
    for i := 0; i < lh; i++ {
        right[i] = lh
    }
    stack := []int {}
    for i := 0; i < lh; i++ {
        ls := len(stack)
        for ls != 0 && heights[stack[ls - 1]] > heights[i] {
            right[stack[ls - 1]] = i
            stack = stack[:ls - 1]
            ls--
        }
        stack = append(stack, i)
    }
    stack = []int {}
    for i := lh - 1; i >= 0; i-- {
        ls := len(stack)
        for ls != 0 && heights[stack[ls - 1]] > heights[i] {
            left[stack[ls - 1]] = i + 1
            stack = stack[:ls - 1]
            ls--
        }
        stack = append(stack, i)
    }
    res := 0
    for i := 0; i < lh; i++ {
        area := (right[i] - left[i]) * heights[i]
        if area > res {
            res = area
        }
    }
    return res
}