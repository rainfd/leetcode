/*
计算某个十进制的数字在二进制上存在多少个1
二进制 十进制 1的数量
0 0 0
1 1 1
10 2 1
11 3 2
100 4 1
101 5 2
110 6 2
111 7 3
1000 8 1
1001 9 2
1010 10 2
1011 11 3
1100 12 2
1101 13 3
1110 14 3
1111 15 4
10000 16 1

F(x) 为十进制数字x在二进制上1的数量

F(x) = 

x = 0 F(0) = 0

if x % 2 == 0 
    F(x) = F(x/2)
if x % 2 == 1
    F(x) = F((x-1)/2) + 1
    
    
*/

func countBits(num int) []int {
  record := []int{}
  for i := 0; i <= num; i++ {
    var x int
    if i == 0 {
      x = 0
    } else if i % 2 == 0 {
      x = record[i/2]
    } else {
      x = record[(i-1)/2] + 1
    }
    record = append(record, x)
  }
  return record
}
