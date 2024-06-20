func maxProfit(prices []int) int {
    maxProfit := 0
    minPrice := prices[0]

    for _, price := range prices {
        minPrice = min(minPrice, price)
        maxProfit = max(maxProfit, price - minPrice)
    }

    return maxProfit
}