func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }
    la, lb := len(a), len(b)
    res := make([]byte, len(a))
    c := byte('0')
    for ib := lb - 1; ib >= 0; ib-- {
        ia := ib + la - lb
        if a[ia] == b[ib] {
            res[ia] = c
            c = a[ia]
        } else if c == '0' {
            res[ia] = '1'
        } else {
            res[ia] = '0'
        }
    }
    for i := la - lb - 1; i >= 0; i-- {
        if c == '1' {
            if a[i] == '1' {
                res[i] = '0'
            } else {
                res[i] = '1'
                c = '0'
            }
        } else {
            res[i] = a[i]
        }
    }
    r := string(res)
    if c == '1' {
        r = "1" + r
    }
    return r
}