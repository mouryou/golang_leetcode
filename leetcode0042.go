func trap(height []int) int {
    res := 0
    l, r := 0, len(height) - 1
    for l < r {
        if height[l] <= height[r] {
            lmax := height[l]
            for l < r && height[l] <= lmax {
                res += lmax - height[l]
                l += 1
            }
        } else {
            rmax := height[r]
            for height[r] <= rmax {
                res += rmax - height[r]
                r -= 1
            }
        }
    }
    return res
}