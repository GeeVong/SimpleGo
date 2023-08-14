package main

import (
	"fmt"
)

const (
	Color int = 4
	Cols  int = 8
	Rows  int = 4
	Step  int = 1
)

var grid [Rows][Cols]int
var maxScore int

func main() {
	// 初始化棋盘
	initGrid()

	// 计算通过一步操作可以得到的最大得分
	maxScore = calculateMaxScore(grid)
	fmt.Println("通过一步操作可以得到的最大得分为：", maxScore)

	// 利用回溯算法找出 step 交换步骤之后的最大得分
	backtrack(grid, Step, 0)
	fmt.Println("通过交换", Step, "次操作可以得到的最大得分为：", maxScore)
}

// 初始化棋盘
func initGrid() {
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			grid[i][j] = i%Color + 1
		}
	}
}

// 计算通过一步操作可以得到的最大得分
func calculateMaxScore(grid [Rows][Cols]int) int {
	maxScore := 0

	// 横向连续消除
	for i := 0; i < Rows; i++ {
		cnt := 1
		for j := 1; j < Cols; j++ {
			if grid[i][j] == grid[i][j-1] {
				cnt++
			} else {
				if cnt >= 3 {
					maxScore += getScore(cnt)
				}
				cnt = 1
			}
		}
		if cnt >= 3 {
			maxScore += getScore(cnt)
		}
	}

	// 纵向连续消除
	for j := 0; j < Cols; j++ {
		cnt := 1
		for i := 1; i < Rows; i++ {
			if grid[i][j] == grid[i-1][j] {
				cnt++
			} else {
				if cnt >= 3 {
					maxScore += getScore(cnt)
				}
				cnt = 1
			}
		}
		if cnt >= 3 {
			maxScore += getScore(cnt)
		}
	}

	return maxScore
}

// 回溯算法找出 step 交换步骤之后的最大得分
func backtrack(grid [Rows][Cols]int, step, curScore int) {
	if step == 0 {
		if curScore > maxScore {
			maxScore = curScore
		}
		return
	}

	// 横向交换相邻元素
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols-1; j++ {
			if grid[i][j] != grid[i][j+1] {
				grid[i][j], grid[i][j+1] = grid[i][j+1], grid[i][j]
				score := calculateMaxScore(grid)
				backtrack(grid, step-1, curScore+score)
				grid[i][j], grid[i][j+1] = grid[i][j+1], grid[i][j]
			}
		}
	}

	// 纵向交换相邻元素
	for j := 0; j < Cols; j++ {
		for i := 0; i < Rows-1; i++ {
			if grid[i][j] != grid[i+1][j] {
				grid[i][j], grid[i+1][j] = grid[i+1][j], grid[i][j]
				score := calculateMaxScore(grid)
				backtrack(grid, step-1, curScore+score)
				grid[i][j], grid[i+1][j] = grid[i+1][j], grid[i][j]
			}
		}
	}
}

// 根据连续个数获取得分
func getScore(cnt int) int {
	switch cnt {
	case 3:
		return 1
	case 4:
		return 4
	case 5:
		return 10
	default:
		return 0
	}
}
