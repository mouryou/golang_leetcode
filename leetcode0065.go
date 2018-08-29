func isNumber(s string) bool {
    s = strings.Trim(s, " ")
    if len(s) == 0 {
        return false
    }
    if s[0] == '+' || s[0] == '-' {
        s = s[1:]
        if len(s) == 0 {
            return false
        }
    }
    dot, e := -1, -1
    for i, r := range s {
        if r == '.' {
            if dot != -1 {
                return false
            } else {
                dot = i
            }
        } else if r == 'e' {
            if e != -1 {
                return false
            } else {
                e = i
            }
        } else if r == '-' || r == '+' {
            if i == 0 || s[i - 1] != 'e' {
                return false
            }
        } else if !unicode.IsDigit(r) {
            return false
        }
    }
    if e != -1 {
        return dot < e && isNumber(s[:e]) && isNumber(s[e + 1:])
    }
    if dot != -1 {
        return len(s) != 1
    }
    return true
}