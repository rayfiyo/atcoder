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
func stringLine() string {
	sc.Scan()
	return sc.Text()
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
func intSplit(N int) []int {
	data := make(map[int]int)
	for i := 0; i < N; i++ {
		data[i] = 0
	}
	sc.Scan()

	// 区切り文字を指定する
	// (sc.Text(), "ここで区切り文字指定だよ～")
	text := strings.Split(sc.Text(), " ")

	// 1文字ずつで区切る
	//text := strings.SplitN(sc.Text(), "", len(sc.Text()))

	ans := []int{}
	for _, t := range text {
		num, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		data[num]++
		if data[num] == 2 {
			ans = append(ans, num)
		}
	}
	return ans
}

// 対話形式の問い は TLEなる！
func main() {
	var N int
	fmt.Scanf("%d\n", &N)
	A := intSplit(N)

	for i, _ := range A {
		fmt.Printf("%d ",A[i])
	}
}
