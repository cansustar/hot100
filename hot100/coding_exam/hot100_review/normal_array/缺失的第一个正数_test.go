package normalarray

import "testing"

/*
 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
 eg:
输入：nums = [1,2,0]
输出：3
解释：范围 [1,2] 中的数字都在数组中。

输入：nums = [3,4,-1,1]
输出：2
解释：1 在数组中，但 2 没有。

*/

func theFirstLostNum(nums []int) int {
	n := len(nums)

	// 第一步：将所有的负数、零、和大于n的数替换为n+1，因为它们不会影响到缺失的第一个正数
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	// 第二步：利用元素的值来作为索引，将数组中存在的正数标记到对应的位置
	for i := 0; i < n; i++ {
		// 因为负数、零、和大于n的数已经被替换为n+1，所以这里不会取到那些已经被处理过的数
		num := abs(nums[i])
		if num <= n {
			// 我们通过将 nums[num-1] 标记为负数来表示 num 存在于数组中
			nums[num-1] = -abs(nums[num-1])
		}
	}

	// 第三步：找出第一个正数的位置，即返回第一个值大于0的索引+1
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	// 如果所有的1到n都存在，那么返回n+1
	return n + 1
}

// 辅助函数：计算绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestFirst(t *testing.T) {
	nums := []int{3, 4, -1, 1}
	result := theFirstLostNum(nums)
	t.Log(result) // 期望输出2
}
