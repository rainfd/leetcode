package main

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		length := len(A[i])
		for j := 0; j < length/2; j++ {
			A[i][j], A[i][length-j-1] = A[i][length-j-1], A[i][j]
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			A[i][j] = 1 - A[i][j]
		}
	}

	return A
}
