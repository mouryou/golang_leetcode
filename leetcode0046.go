func permute(nums []int) [][]int {  
    results := make([][]int, 0)
    search(nums, make([]int, 0), make([]bool, len(nums)), &results)
    return results
}

func search(nums, result []int, chosen []bool, results *[][]int) {
    if len(result) == len(nums) {
        *results = append(*results, append([]int {}, result...)) 
    }
    for i, n := range nums {
        if !chosen[i] {
            result = append(result, n)
            chosen[i] = true
            search(nums, result, chosen, results)
            chosen[i] = false
            result = result[:len(result) - 1]
        }
    }
}