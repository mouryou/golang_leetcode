func maxProfit(prices []int) int {
    l := len(prices)
    if l == 0 {
        return 0
    }
    buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0
    for _, n := range prices {
        if n + buy2 > sell2 {
            sell2 = n + buy2
        }
        if sell1 - n > buy2 {
            buy2 = sell1 - n
        }
        if n + buy1 > sell1 {
            sell1 = n + buy1
        }
        if -n > buy1 {
            buy1 = -n
        }
    }
    return sell2
}