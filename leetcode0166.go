func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }
    sign := 1
    if numerator < 0 {
        sign *= -1
        numerator = -numerator
    }
    if denominator < 0 {
        sign *= -1
        denominator = -denominator
    }
    prefix := ""
    if sign == -1 {
        prefix = "-"
    }
    res := strconv.Itoa(numerator / denominator)
    numerator = numerator % denominator
    if numerator == 0 {
        return prefix + res
    }
    res += "."
    set := make(map[[2]int]int)
    quotients := make([]int, 0)
    i := 0
    for numerator % denominator != 0 {
        k := [2]int {numerator, denominator}
        index, ok := set[k]
        if ok {
            for j := 0; j < index; j++ {
                res += strconv.Itoa(quotients[j])
            }
            res += "("
            for j := index; j < len(quotients); j++ {
                res += strconv.Itoa(quotients[j])
            }
            res += ")"
            return prefix + res
        }
        numerator *= 10
        q := numerator / denominator
        numerator = numerator % denominator
        quotients = append(quotients, q)
        set[k] = i
        i++
    }
    for _, q := range quotients {
        res += strconv.Itoa(q)
    }
    return prefix + res
}