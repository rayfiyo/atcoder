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

	switchON := make([]int, m)
	for i := 0; i < m; i++ {
		if u[i] == 1 {
			switchON[i] = 1
		} else {
			switchON[i] = random(2)
		}
	}

	for i := 0; i < n; i++ {
		if switchON[i] == 1 {
			switch u[i] {
				case x[i]:
					fmt.Printf("%d ", random(3001))
				case y[i]:
					fmt.Printf("%d ", random(3001))
			}
		} else {
			fmt.Printf("0 ")
		}
	}
	fmt.Printf("\n")

	for i := 0; i < m; i++ {
		fmt.Printf("%d ", switchON[i])
	}
	fmt.Printf("\n")
}
