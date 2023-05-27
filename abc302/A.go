/*
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

	A, B := 1, 1
	fmt.Scanf("%d %d\n", &A, &B)

	if B == 1 {
		fmt.Println(A)
	} else if A <= B {
		fmt.Println("1")
	} else if A%B == 0 {
		fmt.Println(A / B)
	} else if A%B != 0 {
		fmt.Println((A / B) + 1)
	} else {
		fmt.Println("omg")
	}
}
*/
