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
	N := 100
	D := 100
	fmt.Scanln(&N, &D)

	S := make([][]string, N)
	for i := 0; i < N; i++ {
		S[i] = make([]string, D)
	}
	for i := 0; i < N; i++ {
		S[i] = scstring()
	}

	okDay := make([]int, D)  // okDay[j] == N なら、j日みんなOK
	for j := 0; j < D; j++ { // j日目の
		for i := 0; i < N; i++ { //iさん
			if S[i][j] == "o" {
				okDay[j]++
			}
		}
	}

	fmt.Println(findAnsDay(okDay, N))

	/*start := make([]int, N)
	end := make([]int, N)

	for i := 0; i < N; i++ {
		for j := 0; j < D; j++ {
			if j == 0 {
				if S[i][j] == "o" {
					start[i] = 0
				} else if S[i][j] == "x" {
					start[i] = -1
				}
			} else if start[i] = 0 {
				if S[i][j] == "x" && end[i] == 0{
					end[i] = j-1
				} else if S[i][j] == "x" {
					start[i] = -1
				}
			}

			if S[i][j] == "o" && start[i] == -1 && end[i] == 0 {
				start[i] = j
			} else if S[i][j] == "x" && start[i] != 0 && end[i] == 0 {
				end[i] = j
			}
		}
	}*/
}
func findAnsDay(okDay []int, D int) int {
	checkDay := 0
	ansDay := 0

	for _, num := range okDay {
		if num == D {
			checkDay++
			if checkDay > ansDay {
				ansDay = checkDay
			}
		} else {
			checkDay = 0
		}
	}

	return ansDay
}

/*func findAnsDay(okDay []int) int {
	ansDay := []int{}
	check := []int{}

	for i, day := range okDay {
		if i == 0 || day == okDay[i-1]+1 { // 最初 か 連続してる 要素
			check = append(check, day)
		} else {
			if len(check) > len(ansDay) {
				ansDay = append([]int{}, check...)
			}
			check = []int{day} // Start a new sequence
		}
	}

	if len(check) > len(ansDay) { // 最後のループ用
		ansDay = append([]int{}, check...)
	}

	return len(ansDay)
}*/

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能
//var sr= bufio.NewReader(os.Stdin)	// 低レイヤ、区切り指定、バイナリ入力可能

// string型スライス 1行入力 指定区切り
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

// int型スライス 1行入力 指定区切り
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

// int型 1行入力
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
	day := []int{5, 2, 9, 1, 7}
	sort.Ints(day)
	fmt.Println(day)

	// - - - //

	dayber := 123.4
	// 四捨五入
	fmt.Println(math.Round(dayber))
	// 切り上げ
	fmt.Println(math.Ceil(dayber))
	// 切り捨て
	fmt.Println(math.Floor(dayber))

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
