func countAndSay(n int) string {
    res := "1"
    for i := 0; i < n - 1; i++ {
        res = next(res)
    }
    return res
}

func next(s string) string {
    r, c := s[0], 1
    res := ""
    for i := 1; i < len(s); i++ {
        if s[i] == r {
            c += 1
        } else {
            res += strconv.Itoa(c) + string(r)
            r = s[i]
            c = 1
        }
    }
    return res + strconv.Itoa(c) + string(r)
}