package matrix

import "testing"

/*
搜索二维矩阵
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
	每行中的整数从左到右按升序排列。
	每列中的整数从上到下按升序排列。
*/

func SearchMatrix(matrix [][]int, target int) bool {
	/*
		可以用类似二分查找的方法。因为矩阵从左到右，从上到下越来越大。 所有左下角和右上角趋近于中间值。
		从右上角开始，如果当前值大于target，那么往左走，如果当前值小于target，那么往下走.
	*/
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// 从右上角开始遍历
	row, col := 0, len(matrix[0])-1
	// 因为是从右上角开始遍历，所以要保证行不超过下界，列不超过左边界（0）
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}

func TestFind(t *testing.T) {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 5
	result := SearchMatrix(matrix, target)
	t.Log(result)
}
