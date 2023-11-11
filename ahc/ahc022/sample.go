/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	y int
	x int
}

type Judge struct{}

func (j *Judge) setTemperature(temperature [][]int) {
	for _, row := range temperature {
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(row)), " "), "[]"))
	}
}

func (j *Judge) measure(i, y, x int) int {
	fmt.Printf("%d %d %d\n", i, y, x)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	v, _ := strconv.Atoi(scanner.Text())
	if v == -1 {
		fmt.Fprintf(os.Stderr, "something went wrong. i=%d y=%d x=%d\n", i, y, x)
		os.Exit(1)
	}
	return v
}

func (j *Judge) answer(estimate []int) {
	fmt.Println("-1 -1 -1")
	for _, e := range estimate {
		fmt.Println(e)
	}
}

type Solver struct {
	L           int
	N           int
	S           int
	LandingPos  []Pos
	JudgeObject *Judge
}

func NewSolver(L, N, S int, landingPos []Pos) *Solver {
	return &Solver{
		L:           L,
		N:           N,
		S:           S,
		LandingPos:  landingPos,
		JudgeObject: &Judge{},
	}
}

func (s *Solver) solve() {
	temperature := s.createTemperature()
	s.JudgeObject.setTemperature(temperature)
	estimate := s.predict(temperature)
	s.JudgeObject.answer(estimate)
}

func (s *Solver) createTemperature() [][]int {
	temperature := make([][]int, s.L)
	for i := range temperature {
		temperature[i] = make([]int, s.L)
	}
	for i, pos := range s.LandingPos {
		temperature[pos.y][pos.x] = i * 10
	}
	return temperature
}

func (s *Solver) predict(temperature [][]int) []int {
	estimate := make([]int, s.N)
	for iIn := 0; iIn < s.N; iIn++ {
		fmt.Printf("# measure i=%d y=0 x=0\n", iIn)

		measuredValue := s.JudgeObject.measure(iIn, 0, 0)
		minDiff := 9999
		for iOut, pos := range s.LandingPos {
			diff := abs(temperature[pos.y][pos.x] - measuredValue)
			if diff < minDiff {
				minDiff = diff
				estimate[iIn] = iOut
			}
		}
	}
	return estimate
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var L, N, S int
	fmt.Scanf("%d %d %d\n", &L, &N, &S)
	landingPos := make([]Pos, N)
	for i := 0; i < N; i++ {
		var y, x int
		fmt.Scanf("%d %d\n", &y, &x)
		landingPos[i] = Pos{y: y, x: x}
	}

	solver := NewSolver(L, N, S, landingPos)
	solver.solve()
}
*/
