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
	N := 50   // N人について
	M := 1225 // M個のデータ
	fmt.Scanln(&N, &M)

	// data[A]B (keyがA, valueがB)
	data := make(map[int]int, M)
	for i := 0; i < M; i++ {
		rawData := scint()
		data[rawData]
	}

	// A == data[0]
	// B == data[1]
	//data:= make([][]int, M)
	//for i := 0; i < M; i++ {
	//	data[i] = make([]int, 2)
	//}
	//for i := 0; i < M; i++ {
	//	data[i] = scint()
	//}

}

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// string 区切り1行
func scstring() []string {
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

// int    区切り1行
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
