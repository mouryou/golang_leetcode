func getPermutation(n int, k int) string {
    result := 0
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        nums[i] = i + 1
    }
    factors := make([]int, n)
    factors[0] = 1
    for i := 1; i < n; i++ {
        factors[i] = i * factors[i - 1]
    }
    k -= 1
    for i := 1; i <= n; i++ {
        j := k / factors[n - i]
        result = result * 10 + nums[j]
        nums = append(nums[:j], nums[j + 1:]...)
        k = k % factors[n - i]
    } 
    return strconv.Itoa(result)
}
