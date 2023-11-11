package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(ceiling int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(ceiling)
}

func main() {
	n, m, k := 100, 300, 5000
	fmt.Scanf("%d %d %d\n", &n, &m, &k)

	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &x[i], &y[i])
	}

	u, v, w := make([]int, m), make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d %d\n", &u[i], &v[i], &w[i])
	}

	a, b := make([]int, k), make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d\n", a[i], b[i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", random(2001))
	}
	fmt.Printf("\n")
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", random(2))
	}
	fmt.Printf("\n")
}

