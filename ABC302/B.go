package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin)

// バッファでstrings型として行読み取り
func stringLine() string {
	sc.Scan()
	return sc.Text()
}

// バッファでint型として行読み取り
func intSC() int {
	sc.Scan()
	inputInt, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return inputInt
}

func oneChar() []string {
	slice := []string{}

	sc.Scan()
	text := sc.Text

	for i, oneC := range text {
		slice[i] = oneC
	}

	return slice
}

// 空白区切り スライス
func intLine() []int {
	slice := []int{}

	sc.Scan()
	text := strings.Split(sc.Text(), " ")

	for _, t := range text {
		converted, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		slice = append(slice, converted)
	}
	return slice
}

// 対話形式の問い は TLEなる！
func main() {
	// want to delimited-space: sc.Split(bufio.ScanWords)

	H, W := 5, 5
	fmt.Scanf("%d %d\n", &H, &W)
	for i := 0; i < H; i++ {
		S[i] := stringLine()
	}
}

/* 検討要素1
in := bufio.NewReader(os.Stdin)
out := bufio.NewWriter(os.Stdout)
defer out.Flush()
var N int
var format string
fmt.Fscan(in, &N)
fmt.Fprintln(out, N, format)
*/

/* 未検証1
func printInts(s []int) {
	p := []string{}
	for _, v := range s {
		p = append(p, strconv.Itoa(v))
	}
	fmt.Println(strings.Join(p, " "))
}
*/

/* 未検証2
func printStrings(s []string) {
	fmt.Println(strings.Join(s, " "))
}
*/
