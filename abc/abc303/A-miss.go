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

// 区切り スライス
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

	N := 100
	fmt.Scanf("%d\n", &N)
	S := stringLine()
	T := stringLine()

	for i := 0; i < N; i++ {
		if S[i] != T[i] {
			if string(S[i]) != "1" && string(T[i]) != "l" {
				if string(T[i]) != "1" && string(S[i]) != "l" {
					if string(S[i]) != "0" && string(T[i]) != "o" {
						if string(T[i]) != "0" && string(S[i]) != "o" {
							fmt.Println("No")
							break
						}
					}
				}
			}
		}
		if i >= N-1 {
			fmt.Println("Yes")
		}
	}
}
*/
