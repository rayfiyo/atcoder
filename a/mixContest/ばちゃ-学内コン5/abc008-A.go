/*
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

// strings型スライス 行ごと
func stringLine() string {
	sc.Scan()
	return sc.Text()
}

// int型スライス？ 行ごと
func intSC() int {
	sc.Scan()
	inputInt, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return inputInt
}

// int型スライス 区切り指定
func intLine() []int {
	slice := []int{}

	sc.Scan()
	text := strings.Split(sc.Text(), " ") // 空白区切り
	// (sc.Text(), "ここが区切り文字だよ～")

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
	// 空白区切りならこれつける: sc.Split(bufio.ScanWords)

	S, T := 1000, 1000
	fmt.Scanf("%d %d\n", &S, &T)

	if S < T {
		fmt.Println(T - S + 1)
	} else if S > T {
		fmt.Println(S - T + 1)
	} else if S == T {
		fmt.Println(1)
	} else {
		fmt.Println("omg")
	}
}
*/
