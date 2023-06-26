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
	b := make([][]int, 30)
	for i := 0; i < 30; i++ {
		b[i] = intSplit()
	}

	K := 0
	klog := make([]string, 0, 10000)
	for i := 0; i < 100000; i++ {

		for bx := 0; bx < 29; bx++ {
			for by := 0; by < bx+1; by++ {
				if by < bx+1 { // 制約
					rank, hx, hy := 464, 0, 0
					nx := bx + 1                    // newX 比較(入れ替え)対象、基準の一つ下の段
					for ny := by; ny < by+2; ny++ { // newY 比較(入れ替え)対象
						if nx < 30 && nx != bx { //制約
							if ny > -1 && ny < 30 { // 制約
								if rank > b[nx][ny] {
									rank = b[nx][ny]
									hx = nx
									hy = ny
								}
							}
						}
					}

					if b[bx][by] > b[hx][hy] { // 下に一番小さいやつあるとき
						b[bx][by], b[hx][hy] = b[hx][hy], b[bx][by]
						K++
						text := fmt.Sprintf("%d %d %d %d", bx, by, hx, hy)
						klog = append(klog, text)
						if K > 10000-10 { // 保険
							i = 100000
							by = bx + 1
							break
						}
					}

				}
			}
		}

	}
	fmt.Println(K)
	for i := range klog {
		fmt.Println(klog[i])
	}
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
