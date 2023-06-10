package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	for _, t := range text {
		converted, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		slice = append(slice, converted)
	}
	return slice
}

// 対話形式の問い は TLEなる！
func main() {
	// 空白区切りならこれつける: sc.Split(bufio.ScanWords)

	p, q := "A", "G"
	pn, qn := 1, 24
	fmt.Scanf("%s %s\n", &p, &q)

	// 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24
	// A        B  C           D  E              F                          G

	switch p {
	case "A":
		pn = 1
	case "B":
		pn = 4
	case "C":
		pn = 5
	case "D":
		pn = 9
	case "E":
		pn = 10
	case "F":
		pn = 15
	case "G":
		pn = 24
	}

	switch q {
	case "A":
		qn = 1
	case "B":
		qn = 4
	case "C":
		qn = 5
	case "D":
		qn = 9
	case "E":
		qn = 10
	case "F":
		qn = 15
	case "G":
		qn = 24
	}

	if pn < qn { // pn >= qn に変換
		pn += qn
		qn = pn - qn
		pn -= qn
	}

	fmt.Println(pn - qn)
}
