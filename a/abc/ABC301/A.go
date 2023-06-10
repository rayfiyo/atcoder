/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 対話形式の問い は TLEなる！
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Acount, Tcount int
	var S string

	fmt.Fscan(in, &N)
	fmt.Fscan(in, &S)

	if N == 1 {
		fmt.Println(S[0:1])
	} else {
		for i := 0; i < N; i++ {
			if S[i:i+1] == "T" {
				Acount++
			} else {
				Tcount++
			}

			if Acount >= N/2 {
				fmt.Println("T")
				break
			} else if Tcount >= N/2 {
				fmt.Println("A")
				break
			}
		}
	}
}
*/
