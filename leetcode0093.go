func restoreIpAddresses(s string) []string {
    return search(s, 4)
}

func search(s string, n int) []string {
    res := []string {}
    l := len(s)
    if l > n * 3 || l < n {
        return res
    }
    if n == 1 {
        num, _ := strconv.Atoi(s)
        if num <= 255 && l == len(strconv.Itoa(num)) {
            return []string {s}
        }
    } else {
        for i := 1; i <= 3 && i < l; i++ {
            num, _ := strconv.Atoi(s[:i])
            if num > 255 || i != len(strconv.Itoa(num)) {
                continue
            }
            for _, s2 := range search(s[i:], n - 1) {
                res = append(res, s[:i] + "." + s2)
            }
        } 
    }
    return res
}