package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// 対話形式の問い は TLEなる！
func main() {
	N, M := 100, 100
	fmt.Scanf("%d %d\n", &N, &M)
	C := stringLine()
	D := stringLine()
	P := intSplit()

	ans := 0
	for _, c := range C {
		for i := 0; i < M; i++ {
			if c == D[i] {
				ans += P[i+1]
				break
			} else if i+1 == M {
				ans += P[0]
			}
		}
	}
	fmt.Println(ans)
}

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// string型(スライス？) 1行入力
func stringLine() []string {
	slice := []string{}
	sc.Scan()
	text := strings.Split(sc.Text(), " ")
	for _, tValue := range text {
		slice = append(slice, tValue)
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

	for _, tValue := range text {
		inted, err := strconv.Atoi(tValue)
		if err != nil {
			panic(err)
		}
		slice = append(slice, inted)
	}
	return slice
}

// - * - * - * - * - * -

func memo() {
	// slice
	// 要素数 = 13, 容量 = 15。容量は省略可能。
	slice := make([]int, 13, 15)
	// sliceの (14-1)番目に14 を (15-1)番目に15 を追加。lenは15になる。
	slice = append(slice, 14, 15)
	// 要素は初期値がついてる、容量はnil(？)でスライスの最大長。過剰にメモリ確保防ぐ。
	// 容量超えたら、容量*2 の容量が更に確保される
	fmt.Println(slice)

	// - - - //

	number := 123.4
	// 四捨五入
	fmt.Println(math.Round(number))
	// 切り上げ
	fmt.Println(math.Ceil(number))
	// 切り捨て
	fmt.Println(math.Floor(number))

	// - - - //

	// 0から29までのランダムな整数を2つ作成
	rand.Seed(time.Now().UnixNano())
	ranNum1 := rand.Intn(30)
	ranNum2 := rand.Intn(30)
	fmt.Println(ranNum1)
	fmt.Println(ranNum2)
}
