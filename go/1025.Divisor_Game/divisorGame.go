/*
F(xi) 为 x的所有因子列表

d(x) = 
    d(x-F(xi)) == false => d(x) true
    只要任意一个因数xi的d(xi)为false，d(x)就可以为true
*/

import "math"

func divisorGame(N int) bool {
    records := []bool{true, false}
    
    for i := 2; i <= N; i++ {
        status := false
        for _, j := range factor(i) {
            if records[i - j] == false {
                status = true
                break
            }
        }
        records = append(records, status)
    }
    return records[N]
}

func factor(N int) []int {
    factors := []int{}
    for i := 1; i <= int(math.Sqrt(float64(N))); i++ {
        if N % i == 0 {
            factors = append(factors, i)
            factors = append(factors, N/i)
        }
    } 
    return factors
}
