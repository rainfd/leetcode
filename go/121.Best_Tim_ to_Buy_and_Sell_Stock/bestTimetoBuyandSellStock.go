func maxProfit(prices []int) int {
    sum := 0
    total := 0
    for i := 1; i < len(prices); i++ {
        total += prices[i] - prices[i-1]
        
        if total > sum {
            sum = total
            print(sum)
        }
        
        if total < 0 {
            total = 0
        }
    }
    return sum
}
