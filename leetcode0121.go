func maxProfit(prices []int) int {
    l := len(prices)
    if l == 0 {
        return 0
    }
    mp, curMin := 0, prices[0]
    for i := 1; i < l; i++ {
        if prices[i] - curMin > mp {
            mp = prices[i] - curMin
        }
        if prices[i] < curMin {
            curMin = prices[i]
        }
    }
    return mp
}