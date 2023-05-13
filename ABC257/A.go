package main

import "fmt"

func main() {
	var N, X int
	fmt.Scanf("%d %d", &N, &X)

	if N >= X {
		fmt.Println(string(rune(65))) // 65 == 'A'
	} else {
		if X%N == 0 {
			fmt.Println(string(rune((X / N) + (64))))
		} else {
			fmt.Println(string(rune((X / N) + (65))))
		}
	}
}
