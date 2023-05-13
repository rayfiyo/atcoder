package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// 対話形式の問い は TLEなる！
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	liner := bufio.NewScanner(os.Stdin)
	defer out.Flush()

	var N int
	var A [100]int
	var input, s string

	fmt.Fscan(in, &N)
	liner.Split(bufio.ScanWords)
	liner.Scan()
	input = liner.Text()
	A[0], _ = strconv.Atoi(strings.Split(input, " "))

	for i:=0; i<N; i++{
		fmt.Fprintln(out, A[i], s)
	}
}
