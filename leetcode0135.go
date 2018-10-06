func candy(ratings []int) int {
    start, peak, res, padding, l := 0, 0, 0, 0, len(ratings)
    ratings = append(ratings, ratings[l - 1])
    increasing := true
    for i := 1; i <= l; i++ {
        cur, prev := ratings[i], ratings[i - 1]
        if cur == prev || (!increasing && cur > prev) {
            if ratings[i - 1] > ratings[start] && start == peak {
                peak = i - 1
            }
            a, b := peak - start, i - 1 - peak
            if a + padding > b {
                a++
            } else {
                b++
            }
            res += (1 + padding + a + padding) * a / 2 + (1 + b) * b / 2
            start, peak = i, i
            if cur > prev {
                padding = 1
            } else {
                padding = 0
            }
            increasing = true
        } else if increasing && cur < prev {
            peak = i - 1
            increasing = false
        } 
    }
    return res
}