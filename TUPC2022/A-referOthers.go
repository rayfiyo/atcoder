package main

import "fmt"

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	array := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		fmt.Scanf("%d", &array[i])
	}

	if N < K {
		fmt.Println("Yes")
	} else {	// N >= K
		for i := K; i <= N && array[i] != i; i++ {
			fmt.Println("No")
			array[0] = 1
			i = N
		}
		if array[0] == 0 {
			fmt.Println("Yes")
		}
	}
}
