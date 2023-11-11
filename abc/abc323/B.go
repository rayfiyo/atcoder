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

func main() {
	N := 100 // 2~100
	fmt.Scanln(&N)

	S := make([][]string, N)
	win := make(map[int]int, N)
	count := 0
	for i := 0; i < N; i++ {
		S[i] = scstring()
		for j := 0; j < N; j++ {
			if S[i][j] == "o" {
				count++
			}
		}
		win[i] = count
		count = 0
	}

	// キーをスライスにコピー
	keys := make([]int, 0)
	for v := range win {
		keys = append(keys, v)
	}

	// スライスを値に基づいて降順にソート
	sort.Slice(keys, func(i, j int) bool {
		return win[keys[i]] > win[keys[j]]
	})

	// ソートされたキーを元にキーと値を表示
	for _, key := range keys {
		fmt.Println(key + 1)
	}
}

// バッファ使う 対話形式 フラッシュ忘れない
// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// string １行取得 任意区切り
func scstring() []string {
	slice := []string{}
	sc.Scan()

	// (sc.Text(), "ここで区切り文字指定だよ～")
	//text := strings.Split(sc.Text(), " ")
	// 1文字ずつで区切る
	text := strings.SplitN(sc.Text(), "", len(sc.Text()))

	for _, tValue := range text {
		slice = append(slice, tValue)
	}
	return slice
}

// int    １行取得 任意区切り
func scint() []int {
	slice := []int{}
	sc.Scan()

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

// int 1行 そのまま入力
func sclineint() int {
	sc.Scan()
	inputInt, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return inputInt
}

// - * - * - * - * - * -

func memo() {
	// 二次元配列A を A[ソート][固定] でソート
	A := [][]int{{4, 6}, {2, 4}, {5, 1}}
	sort.Slice(A, func(i, j int) bool {
		return A[i][0] < A[j][0]
	})
	fmt.Println(A)

	// - - - //

	// 文字列配列を辞書順に並び替え
	str := []string{"banana", "apple", "cherry"}
	sort.Strings(str)
	fmt.Println(str)
	// 配列を昇順に並び替え
	num := []int{5, 2, 9, 1, 7}
	sort.Ints(num)
	fmt.Println(num)

	// - - - //

	// Goのmap(連想配列)
	myMap := make(map[string]int)
	myMap["apple"] = 1
	myMap["banana"] = 2
	fmt.Println(myMap["apple"]) // 1 を出力

	// - - -

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

	// - - - //
	// slice
	// 要素数 = 13, 容量 = 15。容量は省略可能。
	slice := make([]int, 13, 15)
	// sliceの (14-1)番目に14 を (15-1)番目に15 を追加。lenは15になる。
	slice = append(slice, 14, 15)
	// 要素は初期値がついてる、容量はnil(？)でスライスの最大長。過剰にメモリ確保防ぐ。
	// 容量超えたら、容量*2 の容量が更に確保される
	fmt.Println(slice)
}
