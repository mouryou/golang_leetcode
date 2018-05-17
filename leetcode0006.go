func abs(a int) int {
    if a >= 0 {
        return a
    }
    return -a
}

func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    rows := make([][]rune, numRows)
    for i, r := range s {
        t := numRows - 1 - abs((i % (2 * numRows - 2)) - (numRows - 1))
        rows[t] = append(rows[t], r)
    }
    
    result := ""
    for _, rs := range rows {
        result += string(rs)
    }
    return result
    
}