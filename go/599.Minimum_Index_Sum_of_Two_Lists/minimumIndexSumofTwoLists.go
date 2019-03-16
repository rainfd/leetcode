func findRestaurant(list1 []string, list2 []string) []string {
    res1 := make(map[string]int)
    res2 := make(map[string]int)
    for i, res := range list1 {
        res1[res] = i
    }
    for i, res := range list2 {
        res2[res] = i
    }
    
    sum := make(map[string]int)
    for key1, value1 := range res1 {
        if value2, ok := res2[key1]; ok {
            sum[key1] = value1 + value2
        }
    }
    
    result := make([]string, 0)
    min := 10000
    for _, value := range sum {
        if value < min {
            min = value
        }
    }
    for key, value := range sum {
        if value == min {
            result = append(result, key)
        }
    }
    
    return result
}