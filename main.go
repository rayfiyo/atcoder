package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Max. input value: 65536
var sc = bufio.NewScanner(os.Stdin)

func stringLine() string {
	sc.Scan()
	return sc.Text()
}

func intSC() int {
	sc.Scan()
	inputInt, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return inputInt
}

// Space-delimited slices
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

	N := intLine()
	fmt.Println(N[0])
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
