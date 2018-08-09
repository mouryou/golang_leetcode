func divide(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    var i, j int
    if dividend > 0 {
        i = 1
    } else {
        i = -1
        dividend = -dividend
    }
    if divisor > 0 {
        j = 1
    } else {
        j = -1
        divisor = -divisor
    }
    res := 0
    for dividend >= divisor {
        a, b := divisor, 1
        for a << 1 <= dividend {
            a <<= 1
            b <<= 1
        }
        dividend -= a
        res += b
    }
    if i == -1 {
        res = -res
    }
    if j == -1 {
        res = -res
    }
    if res > 2147483647 {
        return 2147483647
    }
    if res < -2147483648 {
        return -2147483648
    }
    return res
}