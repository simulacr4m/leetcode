func maxProfit(prices []int) int {
    gain, loss := 0, 0
    buy := true
    for i := 0; i < len(prices)-1; i++ {
        if buy {
            if prices[i] <= prices[i+1] {
                loss += prices[i]
                buy = false
            }
        } else {
            if prices[i] >= prices[i+1] {
                gain += prices[i]
                buy = true
            }
        }
    }
    if !buy {
        gain += prices[len(prices)-1]
    }
    return gain - loss
}
