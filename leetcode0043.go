func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    l1, l2 := len(num1), len(num2)
    res := ""
    c := 0
    for s := 2; s <= l1 + l2; s++ {
        i := 1
        if s - l2 > 1 {
            i = s - l2
        }
        for ; i <= s - 1 && i <= l1; i++ {
            c += int(num1[l1 - i] - '0') * int(num2[l2 - (s - i)] - '0')
        }
        res = string(c % 10 + '0') + res
        c /= 10
    }
    if c > 0 {
        res = strconv.Itoa(c) + res
    }
    return res
}
