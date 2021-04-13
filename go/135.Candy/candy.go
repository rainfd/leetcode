func candy(ratings []int) int {
    length := len(ratings)
    counter := make([]int, length)

    for i := 0; i < length; i++ {
        counter[i] = 1
    }

    // from left to right
    for i := 1; i < length; i++ {
        if ratings[i] > ratings[i-1] {
            counter[i] = counter[i-1] + 1
        }
    }

    // from right to left
    for i := length - 1; i > 0; i-- {
        if ratings[i] < ratings[i-1] {
            // max(counter[i-1], counter[i] + 1)
            if counter[i-1] < counter[i] + 1 {
                counter[i-1] = counter[i] + 1
            }
        }
    }

    sum := 0
    for i := 0; i < length; i++ {
        sum += counter[i]
    }
    return sum
}