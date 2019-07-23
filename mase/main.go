package main

import (
	"fmt"
	"os"
)

type point struct {
	i int
	j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func main() {
	mase := readMase("mase/mase.in")
	steps := walk(mase, point{0, 0}, point{len(mase) - 1, len(mase[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

func walk(mase [][]int, start, end point) [][]int {
	steps := make([][]int, len(mase))
	for i := range steps {
		steps[i] = make([]int, len(mase[i]))
	}
	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:] //去掉头
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(mase)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func readMase(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	mase := make([][]int, row)
	for i := range mase {
		mase[i] = make([]int, col)
		for j := range mase[i] {
			fmt.Fscanf(file, "%d", &mase[i][j])
		}
	}
	return mase
}
