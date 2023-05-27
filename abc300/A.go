package main

import (
	"fmt"
)

func main() {
	var C [300]int
	N, A, B, sum := 300, 1000, 1000, 2000
	fmt.Scanf("%d %d %d\n", &N, &A, &B)
	sum = A + B
	fmt.Scanf("%d", &C[0])
	for i := 0; i < N; i++ {
		if C[i] == sum {
			fmt.Println(i+1)
			break
		}
		fmt.Scanf("%d", &C[i+1])
	}
}
