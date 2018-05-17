func reverse(x int) int {
    if x == 0 {
        return 0
    }
    var sign int
    if x > 0 {
        sign = 1
    } else {
        sign = -1
        x = -x
    }
    result := 0
    for x > 0 {
        result = result * 10 + (x % 10)
        x /= 10
        if result > (1 << 31) || (result == (1 << 31) && sign != -1) {
            return 0
        } 
    }
    return sign * result
}