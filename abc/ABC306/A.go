/*
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func memo() {
	number := 123.4
	// 四捨五入
    fmt.Println(math.Round(number))
	// 切り上げ
	fmt.Println(math.Ceil(number))
	// 切り捨て
	fmt.Println(math.Floor(number))
}

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// strings型(スライス？) 1行入力
func stringLine() []string {
	slice := []string{}
	sc.Scan()
	text := strings.SplitN(sc.Text(), "", len(sc.Text()))
	for _, t := range text {
		slice = append(slice, t)
	}
	return slice
}

// int型 1行入力
func intSC() int {
	sc.Scan()
	inputInt, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return inputInt
}

// int型スライス 1行入力 指定区切り
func intSplit() []int {
	slice := []int{}
	sc.Scan()

	// 区切り文字を指定する
	// (sc.Text(), "ここで区切り文字指定だよ～")
	text := strings.Split(sc.Text(), " ")

	// 1文字ずつで区切る
	//text := strings.SplitN(sc.Text(), "", len(sc.Text()))

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

	N := 50
	fmt.Scanf("%d\n", &N)
	S := stringLine()
	var ans[100] string
	for i, j:=0, 0; i<N; i++ {
		ans[j] = S[i]
		j++
		ans[j] = S[i]
		j++
	}
	for _, in := range ans {
		fmt.Printf("%s", in)
	}
}
*/
