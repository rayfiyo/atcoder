package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 対話形式の問い は TLEなる！
func main() {
	// 入力
	b := make([][]int, 30)
	for i := 0; i < 30; i++ {
		b[i] = intSplit()
	}

	// 下から比較(実装可能) ∧ 上から揃う(最良)
	bTop := make([][]int, 30)
	for i := range b {
		for j := range b[i] {
			bTop[i] = append(bTop[i], b[i][j])
		}
	}
	kTop := 0
	kTopLog := make([]string, 0, 10000)
	for i := 0; i < 50000; i++ {
		for bx := 28; bx > -1; bx-- { // bx と bx+1(下) の比較なので 0~28
			for by := 0; by < bx+1; by++ { // by と by+1(右) の比較なので 0~bx
				nx := bx + 1 // newX 比較対象: 上の段
				ny := by     // newY 比較対象: 左(基準)
				// nx と ny は定義上、制約内
				if bTop[nx][ny] > bTop[nx][ny+1] { // 右(ny+1)の方が小さいとき
					ny++
				}
				if bTop[bx][by] > bTop[nx][ny] { // 下に一番小さいやつあるとき
					bTop[bx][by], bTop[nx][ny] = bTop[nx][ny], bTop[bx][by]
					text := fmt.Sprintf("%d %d %d %d", bx, by, nx, ny)
					kTopLog = append(kTopLog, text)
					kTop++
				}
			}
		}
	}

	/*
		// 上から比較∧下から揃う
		bLow := make([][]int, 30)
		for i := range b {
			for j := range b[i] {
				bLow[i] = append(bLow[i], b[i][j])
			}
		}
		kLow := 0
		kLowLog := make([]string, 0, 10000)
		for i := 0; i < 50000; i++ {
			for bx := 1; bx < 30; bx++ { // bx と bx-1(上) の比較なので 1~29
				for by := 1; by < bx+1; by++ { // by-1(左) と by の比較なので 1~bx-1
					nx := bx - 1                       // newX 比較対象: 上の段
					ny := by - 1                       // newY 比較対象: 左(基準)
					if bLow[nx][ny] < bLow[nx][ny+1] { // 右(ny+1)の方が大きいとき
						ny++
					}
					if bLow[bx][by] < bLow[nx][ny] { // 上に一番大きいやつあるとき
						bLow[bx][by], bLow[nx][ny] = bLow[nx][ny], bLow[bx][by]
						text := fmt.Sprintf("%d %d %d %d", bx, by, nx, ny)
						kLowLog = append(kLowLog, text)
						kLow++
					}
				}
			}
		}
	*/

	fmt.Println(kTop)
	for i := range kTopLog {
		fmt.Println(kTopLog[i])
	}
	/*
		// 出力
		if kTop < kLow {
			fmt.Println(kTop)
			for i := range kTopLog {
				fmt.Println(kTopLog[i])
			}
		} else {
			fmt.Println(kLow)
			for i := range kLowLog {
				fmt.Println(kLowLog[i])
			}
		}
	*/
}

// 入力最大値: 65536
var sc = bufio.NewScanner(os.Stdin) // 高級、行区切り、トークン分割可能

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
