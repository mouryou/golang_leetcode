func intToRoman(num int) string {
    res := ""
    k := []int {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    v := []string {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    for i := 0; i < len(k); i++ {
        for num >= k[i] {
            res += v[i]
            num -= k[i]
        }
    }
    return res
}