/*
package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 対話形式の問い は TLEなる！
func main() {
	// 入力1
	N := 100
	fmt.Scanln(&N)

	// 入力2
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = intSplit()
	}

	// 出力する配列さくせー
	outline := append([][]int{}, A...)
	for i := 0; i < N; i++ {
		outline[i] = append([]int{}, A[i]...)
	}

	// 処理
	outline[0][0] = A[1][0]
	for i := 0; i < N-1; i++ {
		outline[0][i+1] = A[0][i]
		outline[i+1][N-1] = A[i][N-1]
	}
	for i := N - 1; i > 0; i-- {
		outline[N-1][i-1] = A[N-1][i]
		outline[i-1][0] = A[i][0]
	}

	// 出力
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d", outline[i][j])
		}
		fmt.Println("")
	}
}

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// string型スライス 1行入力 指定区切り
func stringLine() []string {
	slice := []string{}
	sc.Scan()

	// (sc.Text(), "ここで区切り文字指定だよ～")
	text := strings.Split(sc.Text(), " ")
	// 1文字ずつで区切る
	//text := strings.SplitN(sc.Text(), "", len(sc.Text()))

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

	// (sc.Text(), "ここで区切り文字指定だよ～")
	//text := strings.Split(sc.Text(), " ")
	// 1文字ずつで区切る
	text := strings.SplitN(sc.Text(), "", len(sc.Text()))

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
	// 文字列配列を辞書順に並び替え
	str := []string{"banana", "apple", "cherry"}
	sort.Strings(str)
	fmt.Println(str)

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
*/
