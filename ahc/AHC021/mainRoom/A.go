package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"math/rand"
	"time"
)

// 対話形式の問い は TLEなる！
func main() {
	b := make([][]int, 30)
	for i := 0; i < 30; i++ {
		b[i] = intSplit()
	}

	rand.Seed(time.Now().UnixNano())
	K := 0
	klog := make([]string, 0, 10000)
	for i := 0; i < 110000; i++ {

		for bx := 0; bx < 30; bx++ {
			for by := 0; by < bx+1; by++ {
				if by < bx+1 { // 制約
					nx := rand.Intn(2) // newX 変える先
					nx = bx + (nx - 1)
					ny := rand.Intn(2) // newY 変える先
					ny = by + (ny - 1)

					if nx > -1 && nx < 30 && bx != nx { //制約
						if ny > -1 && ny < 30 && ny < nx+1 { // 制約
							if b[bx][by] < b[nx][ny] && bx > nx {
								b[bx][by], b[nx][ny] = b[nx][ny], b[bx][by]
								K++
								text := fmt.Sprintf("%d %d %d %d", bx, by, nx, ny)
								klog = append(klog, text)
							}
						}
					}
				}
			}
		}

		if K > 10000-1 { // 保険
			break
		}

	}
	fmt.Println(K)
	for i := range klog {
		fmt.Println(klog[i])
	}
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

	//---

	number := 123.4
	// 四捨五入
	fmt.Println(math.Round(number))
	// 切り上げ
	fmt.Println(math.Ceil(number))
	// 切り捨て
	fmt.Println(math.Floor(number))
}
