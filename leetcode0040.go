func combinationSum2(candidates []int, target int) [][]int {
    sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
    // dp:
    // key1: startIndex int
    // key2: curTarget int
    // value: result of combinationSum2(candidates[startIndex:], curTarget)
    return search(candidates, target, 0, make(map[int]map[int][][]int))
}

func search(candidates []int, curTarget int, startIndex int, dp map[int]map[int][][]int) [][]int {
    if curTarget == 0 {
        return [][]int {[]int {}}
    }
    l := len(candidates)
    if startIndex >= l || curTarget < candidates[l - 1] {
        return [][]int {}
    }
    if _, ok := dp[startIndex]; ok {
        if results, ok := dp[startIndex][curTarget]; ok {
            return results
        }
    } else {
        dp[startIndex] = make(map[int][][]int)
    }
    results := make([][]int, 0)
    for i := startIndex; i < len(candidates); i++ {
        if curTarget >= candidates[i] && (i == startIndex || candidates[i] != candidates[i - 1]) {
            for _, result := range search(candidates, curTarget - candidates[i], i + 1, dp) {
                duplicateResult := make([]int, len(result))
                copy(duplicateResult, result)
                results = append(results, append(duplicateResult, candidates[i]))
            }
        }
    }
    dp[startIndex][curTarget] = results
    return results
}