package double_pointer

import (
	"math"
	"testing"
)

/*
接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func RainWater(nums []int) int {
	/*
			可以按列计算，计算每一列能接多少雨水，
			即计算每一列左右两边`最高`的柱子，然后取两者中较小的柱子减去当前柱子的高度，即为该列能接的最大数。
			可以用双指针，分别指向数组的两端，作为启始时的两边最大高度
			然后移动指针，更新二者的最大值。并计算指针指向的列能接的雨水数直到两个指针相遇。
		时间复杂度O(n), 空间复杂度O(1)
	*/
	ans := 0
	// 初始化左右指针，可以把left和right就看作当前指向的需要计算的列
	left, right := 0, len(nums)-1
	// 初始化左右两边的最大高度
	leftMax, rightMax := 0, 0
	// 终止条件为左右指针相遇
	for left < right {
		// 分别计算左右两边的最大高度
		leftMax = int(math.Max(float64(leftMax), float64(nums[left])))
		rightMax = int(math.Max(float64(rightMax), float64(nums[right])))
		// 当leftMax较小时，计算左指针指向的列能接的雨水数，反之亦然
		if leftMax < rightMax {
			ans += leftMax - nums[left]
			// 移动left指针
			left++
		} else {
			ans += rightMax - nums[right]
			right--
		}
	}
	return ans
}

func TestRainWater(t *testing.T) {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := RainWater(nums)
	t.Log(result)
}
