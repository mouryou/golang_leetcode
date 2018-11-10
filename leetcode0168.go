func convertToTitle(n int) string {
    res := ""
    for n > 0 {
        res = string('A' + (n - 1) % 26) + res
        n = (n - 1) / 26
    }
    return res
}