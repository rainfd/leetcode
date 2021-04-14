package main

func lemonadeChange(bills []int) bool {
    counter :=  map[int]int{
        5: 0,
        10: 0,
        20: 0,
    }
    for _, bill := range bills {
        if bill == 5 {
            counter[5] += 1
        } else if bill == 10 {
            if counter[5] == 0 {
                return false
            }
            counter[5] -= 1
        } else if bill == 20 {
            if counter[10] > 0 && counter[5] > 0 {
                counter[10] -= 1
                counter[5] -= 1
            } else if counter[5] > 3 {
                counter[5] -= 3
            } else {
                return false
            }
        }
    }
    return true
}
