func maxSubArray(nums []int) int {
    sum := math.MinInt32
    total := 0
    for i := 0; i < len(nums); i++ {
        total += nums[i]
        if total > sum {
            sum = total
        }
        if total < 0 {
            total = 0
        }
    }
    return sum
}
