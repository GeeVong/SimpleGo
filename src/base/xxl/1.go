//package main.go.go
//
//import "fmt"
//
////var (
////	board     [][]int
////	rows      int
////	cols      int
////	color     int
////	maxScores int
////)
//
//func eliminateAndFill1(row, col int) int {
//	if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != color {
//		return 0
//	}
//
//	count := 1
//	board[row][col] = 0
//
//	// 当前方块的上下左右四个方向
//	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
//
//	for _, d := range directions {
//		newRow, newCol := row+d[0], col+d[1]
//		count += eliminateAndFill1(newRow, newCol)
//	}
//
//	return count
//}
//
//func getMaxScoresAfterOneStep(color, cols, rows int) int {
//	maxScores := 0
//
//	// 棋盘初始化
//	board = make([][]int, rows)
//	for i := 0; i < rows; i++ {
//		board[i] = make([]int, cols)
//	}
//
//	// 计算初始得分
//	for row := 0; row < rows; row++ {
//		for col := 0; col < cols; col++ {
//			count := eliminateAndFill(row, col)
//
//			if count >= 3 {
//				switch count {
//				case 3:
//					maxScores += 1
//				case 4:
//					maxScores += 4
//				default:
//					maxScores += 10
//				}
//			}
//		}
//	}
//
//	return maxScores
//}
//
//func backtrack(step, maxSteps, prevScores int) {
//	if step > maxSteps {
//		return
//	}
//
//	for row := 0; row < rows-1; row++ {
//		for col := 0; col < cols-1; col++ {
//			if board[row][col] != board[row][col+1] {
//				// 交换相邻两个方块
//				board[row][col], board[row][col+1] = board[row][col+1], board[row][col]
//
//				// 消除并填充
//				scores := getMaxScoresAfterOneStep(color, cols, rows)
//
//				if scores > maxScores {
//					maxScores = scores
//				}
//
//				backtrack(step+1, maxSteps, scores)
//
//				// 回溯，恢复交换前的方块
//				board[row][col], board[row][col+1] = board[row][col+1], board[row][col]
//			}
//		}
//	}
//}
//
//func main1() {
//	color = 4
//	cols = 8
//	rows = 4
//
//	// 棋盘初始化
//	board = make([][]int, rows)
//	for i := 0; i < rows; i++ {
//		board[i] = make([]int, cols)
//	}
//
//	fmt.Println("Max scores after one step:", getMaxScoresAfterOneStep(color, cols, rows))
//
//	maxSteps := 2 // 设置最大交换次数，可根据需要调整
//	backtrack(1, maxSteps, 0)
//
//	fmt.Println("Max scores after two steps:", maxScores)
//}
