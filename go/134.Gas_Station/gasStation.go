func canCompleteCircuit(gas []int, cost []int) int {
    length := len(gas)
    if length == 0 {
        return -1
    }

    get := []int{}
    for i := 0; i < length; i++ {
        g := gas[i] - cost[i]
        get = append(get, g)
    }

    sum := 0
    for i := 0; i < length; i++ {
        sum += get[i]
    }
    if sum < 0 {
        return -1
    }

    start, sign := 0, 0
    total, max := 0, 0
    for i := 0; i < length * 2 - 1; i++ {
        // queue add
        if start - i >= length {
            total = total + get[i%length] - get[start]
            start += 1
        } else {
            total += get[i%length]
        }

        if total < 0 {
            total = 0
            start = i + 1
        }
        if total > max {
            sign = start
            max = total
        }
    }
    return sign
}
