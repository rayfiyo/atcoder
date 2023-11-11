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

func main() {
	N := 100
	fmt.Scanln(&N)
	S := scstring()

check:
	for i := 0; i < N-1; i++ {
		switch S[i] {
		case "a":
			if S[i+1] == "b" {
				fmt.Println("Yes")
				break check
			}
		case "b":
			if S[i+1] == "a" {
				fmt.Println("Yes")
				break check
			}
		}
		if i+1 >= N-1 {
			fmt.Println("No")
		}
	}

}

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// string 1行取得 任意区切り
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

// int    1行取得 任意区切り
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

// mapの値に基づき並び替えたとき の キーの配列
func sortmap(N map[int]int) []int {
	keys := make([]int, 0, len(N))
	for vofN := range N {
		keys = append(keys, vofN)
	}
	sort.Slice(keys, func(i, j int) bool {
		if N[keys[i]] != N[keys[j]] {
			return N[keys[i]] < N[keys[j]] // 値の昇順↑に基づく
			// return N[keys[i]] > N[keys[j]] // 値の降順↓に基づく
		} else {
			return keys[i] < keys[j] // 同値のときキーの昇順↑に基づく
			// return keys[i] > keys[j] // 同値のときキーの降順↓に基づく
		}
	})
	return keys
}

// int 1行取得 そのまま入力 ？？？
// 使わない６溜まったら消す: ABC323,
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
	rand.NewSource(time.Now().UnixNano())
	ranNum1 := rand.Intn(30)
	ranNum2 := rand.Intn(30)
	fmt.Println(ranNum1)
	fmt.Println(ranNum2)

	// - - - //

	// スライス
	// 要素数 = 13, 容量 = 15。容量は省略可能。
	slice := make([]int, 13, 15)
	// スライスの (14-1)番目に14 を (15-1)番目に15 を追加。lenは15になる。
	slice = append(slice, 14, 15)
	// 要素は初期値がついてる、容量はnil(？)でスライスの最大長。過剰にメモリ確保防ぐ。
	// 容量超えたら、容量*2 の容量が更に確保される
	fmt.Println(slice)

	// - - - //

	// mapの値を降順↓に並び替え，キーを出力
	// mapのキーを配列にすることができればmapの比較ができるので，それが真のときにキーを並び替えれば良い
	// 3 2 1 4 (2 1 0 3 に 1加えたもの) の出力を期待
	N := map[int]int{0: 3, 1: 3, 2: 7, 3: 1}
	// mapのキーをkeysスライスに複製
	keys := make([]int, 0, len(N))
	for vofN := range N { // Nのキーを一時的にもつvofN
		keys = append(keys, vofN)
	}
	// keysをNの式が真のときに並び替える
	sort.Slice(keys, func(i, j int) bool { // Nに基づき並び替え
		if N[keys[i]] != N[keys[j]] {
			return N[keys[i]] > N[keys[j]] // 値の降順に基づく
		} else {
			return keys[i] < keys[j] // 同値のときキーの昇順に基づく
		}
	})
	for _, v := range keys { // 出力
		fmt.Println(v + 1)
	}
}

//*/
