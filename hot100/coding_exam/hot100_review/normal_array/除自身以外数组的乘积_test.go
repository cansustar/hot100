package normalarray

import (
	"fmt"
	"testing"
)

/*
除自身以外数组的乘积
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
不要使用除法，且在 O(n) 时间复杂度内完成此题。
eg:
 	输入：nums = [1,2,3,4]
 	输出：[24,12,8,6]
*/

func powerfulProductExceptSelf(nums []int) []int {
	// 进阶做法：在计算完左边的乘积后，直接在计算右边的乘积的同时，将结果存入answer数组中
	// 这样就可以在O(n)的时间复杂度内完成，并且不需要额外存储另一缀，空间复杂度为O(1)
	right := make([]int, len(nums))
	right[len(nums)-1] = 1
	// 求右边的乘积. 因为已经初始化了right[len(nums)-1] = 1，所以从len(nums)-2开始遍历；
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	// 初始化answer数组
	answer := make([]int, len(nums))
	pre := 1
	for i := 0; i < len(nums); i++ {
		answer[i] = right[i] * pre
		pre = pre * nums[i]
	}
	return answer
}

func productExceptSelf(nums []int) []int {
	// 可以使用前后缀，分别存储元素i左边所有元素的乘积和右边所有元素的乘积
	// 普通做法：两个for循环，一个for循环求左边的乘积，一个for循环求右边的乘积。这样时间复杂度是O(n),空间复杂度是O(n)
	answer := make([]int, len(nums))
	// 先求左边的乘积
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1
	// 求左边的乘积. 因为已经初始化left[0] = 1，所以从1开始遍历；
	//因为left[i] = left[i-1]*nums[i-1]，所以遍历到len(nums)-1
	for i := 1; i <= len(nums)-1; i++ {
		// 索引为i的元素的左边所有元素的乘积，等于索引为i-1的元素的左边所有元素的乘积乘以索引为i-1的元素
		left[i] = left[i-1] * nums[i-1]
		fmt.Println(left)
	}
	// 求右边的乘积. 因为已经初始化了right[len(nums)-1] = 1，所以从len(nums)-2开始遍历；
	//因为right[i] = right[i+1]*nums[i+1]，所以遍历到0
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
		fmt.Println(right)
	}
	// 将左右两边的乘积相乘，就是除自身以外数组的乘积
	for i := 0; i < len(nums); i++ {
		answer[i] = left[i] * right[i]
	}
	return answer
}

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	//result := productExceptSelf(nums)
	result := powerfulProductExceptSelf(nums)
	t.Log(result)
}
