func canPlaceFlowers(flowerbed []int, n int) bool {
	total := 0
	beds := append(flowerbed, 0)

	p1, p2 := 0, 1
	for _, bed := range beds {
		if bed == 0 && p1 == 0 && p2 == 0 {
			// p2 p1 bed
			// 0 0 0    set p1 = 1
			p1 = 1
			total += 1
		}
		if total == n {
			return true
		}
		p2 = p1
		p1 = bed
	}

	return false
}