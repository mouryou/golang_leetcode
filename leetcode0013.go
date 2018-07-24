func romanToInt(s string) int {
    p := 0
    res := 0
    symbol := []string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    value := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    for i := 0; i < len(symbol); i++ {
        for p + len(symbol[i]) <= len(s) && s[p:p + len(symbol[i])] == symbol[i] {
            p += len(symbol[i])
            res += value[i]
        }
    }
    return res
}