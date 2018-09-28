func partition(s string) [][]string {
    palindromes := make(map[int][]int)
    l := len(s)
    for i := 0; i < l; i++ {
        palindromes[i] = append(palindromes[i], i + 1)
        for d := 1; i - d >= 0 && i + d < l; d++ {
            if s[i - d] == s[i + d] {
                palindromes[i - d] = append(palindromes[i - d], i + d + 1)
            } else {
                break
            }
        }
        for d := 1; i + 1 - d >= 0 && i + d < l; d++ {
            if s[i + 1 - d] == s[i + d] {
                palindromes[i + 1 - d] = append(palindromes[i + 1 - d], i + d + 1)
            } else {
                break
            }
        }
    }
    results := make([][]string, 0)
    dfs(0, s, []string {}, palindromes, &results)
    return results
}

func dfs(position int, s string, cur []string, palindromes map[int][]int, results *[][]string) {
    if position == len(s) {
        *results = append(*results, append([]string {}, cur...))
    }
    for _, np := range palindromes[position] {
        cur = append(cur, s[position:np])
        dfs(np, s, cur, palindromes, results)
        cur = cur[:len(cur) - 1]
    }
}
