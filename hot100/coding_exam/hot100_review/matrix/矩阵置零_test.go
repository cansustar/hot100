package matrix

import "testing"

/*
矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。使用原地算法
eg:
 	输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
	输出：[1,0,1],[0,0,0],[1,0,1]】
*/

func MatrixSetZero(matrix [][]int) [][]int {
	// 用两个标记变量，分别标记第一行和第一列是否有0
	row, col := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	// 遍历矩阵，如果matrix[i][j] == 0，则将matrix[i][0]和matrix[0][j]置为0
	for i, r := range matrix {
		for j, v := range r {
			// v就是该单元格的值，如果为0，则将该行和该列的第一个元素置为0
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	// 再次遍历矩阵，如果matrix[i][0]或者matrix[0][j]为0，则将matrix[i][j]置为0
	for i, r := range matrix {
		for j, _ := range r {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func TestMatrix(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	result := MatrixSetZero(matrix)
	t.Log(result)
}
