func min(i, j int) int {
    if i < j {
        return i
    } else {
        return j
    }
}

func maxArea(height []int) int {
    l, r := 0, len(height) - 1
    mv := min(height[l], height[r]) * (r - l)
    for l < r {
        if height[l] < height[r] {
            oh := height[l]
            for l < r {
                l += 1
                if height[l] > oh && min(height[l], height[r]) * (r - l) > mv {
                    mv = min(height[l], height[r]) * (r - l)
                    break
                }
                if height[l] > height[r] {
                    break
                }
            }
        } else {
            oh := height[r]
            for l < r {
                r -= 1
                if height[r] > oh && min(height[l], height[r]) * (r - l) > mv {
                    mv = min(height[l], height[r]) * (r - l)
                    break
                }
                if height[r] > height[l] {
                    break
                }
            }
        }
    }
    return mv
}