func candy(ratings []int) int {
    start, peak, res, l := 0, 0, 0, len(ratings)
    increasing, padding := true, false
    for i := 1; i < l; i++ {
        cur, prev := ratings[i], ratings[i - 1]
        if cur == prev || (!increasing && cur > prev) {
            if ratings[i - 1] > ratings[start] && start == peak {
                peak = i - 1
            }
            a, b := peak - start, i - 1 - peak
            if padding {
                if a + 1 > b {
                    a++
                } else {
                    b++
                }
                res += (2 + a + 1) * a / 2 + (1 + b) * b / 2
            } else {
                if a > b {
                    a++
                } else {
                    b++
                }
                res += (1 + a) * a / 2 + (1 + b) * b / 2
            }
            start, peak = i, i
            padding = cur > prev
            increasing = true
        } else if increasing && cur < prev {
            peak = i - 1
            increasing = false
        } 
    }
    if ratings[l - 1] > ratings[start] && start == peak {
        peak = l - 1
    }
    a, b := peak - start, l - 1 - peak
    if padding {
        if a + 1 > b {
            a++
        } else {
            b++
        }
        res += (2 + a + 1) * a / 2 + (1 + b) * b / 2
    } else {
        if a > b {
            a++
        } else {
            b++
        }
        res += (1 + a) * a / 2 + (1 + b) * b / 2
    }
    return res
}