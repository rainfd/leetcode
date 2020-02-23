func twoSum(nums []int, target int) []int {
    records := make(map[int]int)
    for i, n := range nums {
        records[n] = i
    }
    for i, n := range nums {
        if value, ok := records[target-n]; ok && value != i{
            return []int{i, value}
        }
    }
    return nil
}
