package matrix

import "testing"

/*
 螺旋矩阵
 给你一个m行n列的矩阵matrix，请按照顺时针螺旋顺序，返回矩阵中的所有元素。
eg:
	输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
	输出：[1,2,3,6,9,8,7,4,5]
*/

func MatrixCircle(matrix [][]int) []int {
	// 定义四个边界，分别是上下左右
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	result := make([]int, 0)
	for {
		// 从左到右
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		// 从上到下
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// 从右到左
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}
		// 从下到上
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return result
}

func TestName(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := MatrixCircle(matrix)
	t.Log(result)
}
