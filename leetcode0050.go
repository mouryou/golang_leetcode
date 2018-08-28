func myPow(x float64, n int) float64 {
    if x == 0.0 || x == 1.0 {
        return x
    }
    res := 1.0
    sign := 1
    if n < 0 {
        n = -n
        sign = -1
    }
    for n > 0 {
        if n & 1 == 1 {
            res *= x
        }
        x *= x
        n >>= 1
    }
    if sign == -1 {
        res = 1 / res
    }
    return res
}