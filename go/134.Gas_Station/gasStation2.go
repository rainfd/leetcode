// https://leetcode-cn.com/problems/gas-station/solution/shi-yong-tu-de-si-xiang-fen-xi-gai-wen-ti-by-cyayc/

func canCompleteCircuit(gas []int, cost []int) int {
    minSum := 0
    sum := 0
    sign := 0
    for i := range gas {
        sum += gas[i] - cost[i]
        if sum <= minSum {
            sign = i
            minSum = sum
        }
    }
    if sum < 0 {
        return -1
    }
    return (sign + 1) % len(gas)
}

