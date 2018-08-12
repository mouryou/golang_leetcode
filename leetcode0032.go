func longestValidParentheses(s string) int {
    sc := make([]rune, 0) // stack of character
    si := make([]int, 0) // stack of index
    for i, c := range s {
        if c == ')' && len(sc) > 0 && sc[len(sc) - 1] == '(' {
            sc = sc[:len(sc) - 1]
            si = si[:len(si) - 1]
        } else {
            sc = append(sc, c)
            si = append(si, i)
        }
    }

    if len(si) > 0 {
        maxl := si[0]
        for i := 1; i < len(si); i++ {
            if si[i] - si[i - 1] - 1 > maxl {
                maxl = si[i] - si[i - 1] - 1
            }
        }
        if len(s) - si[len(si) - 1] - 1 > maxl {
            maxl = len(s) - si[len(si) - 1] - 1
        }
        return maxl
    } else {
        return len(s)
    }
    
}