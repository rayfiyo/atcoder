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
	liner := bufio.NewScanner(os.Stdin)
	defer out.Flush()

	var N int
	var S, format string

	fmt.Fscan(in, &N)

	liner.Scan()		// "1 2 3" などの入力を一行スキャン
	S = liner.Text()

	fmt.Fprintln(out, N, format)
	fmt.Fprintln(out, S, format)
}
