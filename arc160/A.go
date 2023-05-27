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

	N, K := 1, 1
	fmt.Scanf("%d %d\n", &N, &K)

	var A []int
	var tmp []int
	//var ans[] int
	A = intLine()
	tmp = A

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			tmp[j] = A[i]
			tmp[i] = A[j]
		}
	}
	fmt.Println(A[0])
	fmt.Println(tmp[0])
}
