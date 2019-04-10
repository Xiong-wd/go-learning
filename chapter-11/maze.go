package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// 从文件中读取数据，首先读取第一列，获取迷宫的大小
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	// 构造迷宫， 获取迷宫的每个数据
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
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

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	// 判断的点的队列，初始值就是起点的队列
	Q := []point{start}
	//判断队列是否为空，只有当队列不为空的时候才继续探索，如果为空，则返回
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 满足三种状况，如果next下一个等于0，
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			// next，在steps中必须没有走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			// 下一个不能是起点
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("chapter-11/maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	// TODO: construct path from steps
}
