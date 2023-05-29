package main

import (
	"fmt"
)

func main() {
	var N, K int

	fmt.Scanf("%d %d", &N, &K)

	array := make([]int, N+1)

	for i := 1; i < N+1; i++ {
		fmt.Scanf("%d", &array[i])
	}

	if N < K {
		fmt.Println("Yes")
	} else { // N >= K
		if array[N] == N && array[K] == K {
			if N == K+1 {
				fmt.Println("Yes")
			} else {
				for i := N - 1; i > K+1; i-- {
					if array[i] != i {
						fmt.Println("No")
						i = K + 1
					}
					if i == K+2 {
						fmt.Println("Yes")
					}
				}
			}
		} else {
			fmt.Println("No")
		}
	}
}
