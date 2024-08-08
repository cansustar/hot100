package normalarray

import "testing"

/*
最大子数和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组
是数组中的一个连续部分。
eg:
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
*/

func MaxSubArray(nums []int) int {
	/*
			常规解法，如果加上当前元素之后的和比当前元素还小，那么就从当前元素开始重新计算和
		时间复杂度O(n), 空间复杂度O(1)
	*/
	// 初始化最大和为0，当前和为0
	maxn, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		// 当前和加上当前元素
		sum += nums[i]
		// 如果大于最大和，就更新最大和，否则最大和不变
		if sum > maxn {
			maxn = sum
		}
		// 如果当前和小于0，就从当前元素开始重新计算和
		if sum < 0 {
			sum = 0
		}
	}
	return maxn
}

func TestMaxSubArray(T *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := MaxSubArray(nums)
	println(result)
}
