func maxProfit(prices []int) int {
    l := len(prices)
    res := 0
    for i := 1; i < l; i++ {
        if prices[i] > prices[i - 1] {
            res += prices[i] - prices[i - 1]
        }
    }
    return res
}