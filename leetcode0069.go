func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    l, r := 0, x / 2 + 1
    for l < r {
        m := (l + r) / 2
        s := m * m
        if s > x {
            r = m
        } else if s < x {
            l = m + 1
        } else {
            return m
        }
    }
    return l - 1
}