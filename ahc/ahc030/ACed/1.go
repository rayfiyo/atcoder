package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read prior information
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])
	fields := make([][][]int, M)
	for i := 0; i < M; i++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		parts = strings.Fields(line)
		ps := make([][]int, len(parts)/2)
		for j := 0; j < len(parts)/2; j++ {
			x, _ := strconv.Atoi(parts[2*j+1])
			y, _ := strconv.Atoi(parts[2*j+2])
			ps[j] = []int{x, y}
		}
		fields[i] = ps
	}

	// drill every square
	hasOil := make([][]int, 0)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("q 1 %d %d\n", i, j)
			resp, _ := reader.ReadString('\n')
			resp = strings.TrimSpace(resp)
			if resp != "0" {
				hasOil = append(hasOil, []int{i, j})
			}
		}
	}

	fmt.Printf("a %d ", len(hasOil))
	for _, point := range hasOil {
		fmt.Printf("%d %d ", point[0], point[1])
	}
	fmt.Println()
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(resp)
	if resp != "1" {
		panic("Unexpected response")
	}
}

