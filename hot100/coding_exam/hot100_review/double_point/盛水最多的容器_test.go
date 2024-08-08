package double_pointer

import (
	"math"
	"testing"
)

/*
盛水最多的容器
给定一个长度为n的整数数组height,有n条垂线，第i条线的两个端点是(i,0)和(i, height[0])
要求找出其中的两条，使其与x轴共同构成的容器可以容纳最多的水。 返回可容纳的最大水量（面积）
*/

// double_point

func MaxArea(height []int) int {
	// 计算面积，即要求长乘高最大。 高则是两条垂线中较短的那一条。
	// 所以可以要求长最大，再移动垂线，找到使得面积最大的位置
	// 故使用左右指针，分别位于两侧
	// 时间复杂度：O(n), 空间复杂度O(1)
	left, right := 0, len(height)-1
	maxArea := 0
	// 终止条件为左右指针相遇
	for left < right {
		// 计算当前的面积，长为两垂线的距离，高为二者中较低的那一个
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		// 更新最大面积
		if area > maxArea {
			maxArea = area
		}
		// 移动两指针中较矮的那一个，尝试寻找更大的面积
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func TestMaxArea(t *testing.T) {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := MaxArea(nums)
	t.Log(result)
}
