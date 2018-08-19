import "sort"

func combinationSum(candidates []int, target int) [][]int {
    sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
    // key1: startIndex 
    // key2: curTarget 
    // value: combinationSum(candidates[startIndex:], curTarget)
    dp := make(map[int]map[int][][]int) 
    return search(candidates, target, 0, dp)
}


func search(candidates []int, curTarget int, startIndex int, dp map[int]map[int][][]int) [][]int {
    if curTarget == 0 {
        return [][]int {[]int {}}
    }
    if curTarget < candidates[len(candidates) - 1] {
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
        if candidates[i] <= curTarget {
            for _, result := range search(candidates, curTarget - candidates[i], i, dp) {
                duplicateResult := make([]int, len(result))
                copy(duplicateResult, result)
                results = append(results, append(duplicateResult, candidates[i]))
            }
            
        }
    }
    dp[startIndex][curTarget] = results
    return results
}