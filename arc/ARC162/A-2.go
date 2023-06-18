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
func intSplit() ([]int, int) {
	slice := []int{}
	sc.Scan()

	// 区切り文字を指定する
	// (sc.Text(), "ここで区切り文字指定だよ～")
	text := strings.Split(sc.Text(), " ")

	// 1文字ずつで区切る
	//text := strings.SplitN(sc.Text(), "", len(sc.Text()))

	c, stop := 1, 0
	for i, tValue := range text {
		inted, err := strconv.Atoi(tValue)
		if err != nil {
			panic(err)
		}
		slice = append(slice, inted)

		if i > 0 && stop < 1 {
			if slice[i-1] == inted-1 {
				c++
			} else {
				stop++
			}
		}
	}
	if c > 1 { // 並んでる
		return slice, c
	} else {
		return slice, 0
	}
}

func noSort(P []int) int {
	if len(P) == 1 {
		return 1
	}
	ans := 0
	for i, v := range P {
		if i+1 > v {
			ans++
		}
	}
	return ans
}

// 対話形式の問い は TLEなる！
func main() {
	T := 500
	fmt.Scanln(&T)
	var N, ans [500]int
	//var P [500][1000]int
	for i := 0; i < T; i++ {
		N[i] = intSC()
		tmp, noSort0 := intSplit()
		if noSort0 == 0 {
			ans[i] = noSort(tmp)
		} else {
			ans[i] = noSort0
		}
		//for j := 0; j < N[i]; j++ {
		//	P[i][j] = tmp[j]
		//}
	}

	//for i := 0; i < T; i++ {
	//	ans[i] = noSort(P[i])
	//}
	for i := 0; i < T; i++ {
		fmt.Println(ans[i])
	}
}
