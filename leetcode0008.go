func myAtoi(str string) int {
    upperbound := (1 << 31) - 1 // the upperbound of overflow
    lowerbound := -(1 << 31) // the lowerbound of overflow
    sign := 0 // the sign of the number, 1 for + and -1 for -
    result := 0 
    started := false // if the number has started
    for _, r := range str {
        if unicode.IsDigit(r) {
            started = true
            if sign == 0 {
                sign = 1
            }
            result = result * 10 + int(r) - int('0')
            if sign * result > upperbound {
                return upperbound
            }
            if sign * result < lowerbound {
                return lowerbound
            }
        } else {
            if started {
                break
            }
            if r == ' ' {
                continue
            } else if r == '+' {
                if sign != 0 {
                    return 0
                }
                sign = 1
                started = true
            } else if r == '-' {
                if sign != 0 {
                    return 0
                }
                sign = -1
                started = true
            } else {
                break
            }
        } 
        
        
    }
    return sign * result
}