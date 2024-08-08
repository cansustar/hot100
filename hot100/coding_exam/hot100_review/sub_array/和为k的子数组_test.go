package subarray

/*
和为k的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。
eg:
输入：nums = [1,1,1], k = 2  因为[1,1 ] 和[ 1,1]
*/

func SubarraySum(nums []int, k int) int {
	/*
		因为要统计的是和为k的子数组的“数量”，所以可以想到用前缀和。
		key是前缀和，value是该前缀和出现的次数。
	*/
	res := 0
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		// 表示arr[0]到arr[i-1]的和
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	// 前缀和的哈希表
	hashMap := make(map[int]int)
	// 初始化前缀和为0的个数为1
	hashMap[0] = 1
	for i := 1; i <= len(nums); i++ {
		// 求前缀和
		preSum := prefixSum[i]
		// 求前缀和为preSum-k的个数
		if v, ok := hashMap[preSum-k]; ok {
			res += v
		}
		// 更新前缀和的个数
		hashMap[preSum]++
	}
	return res
}
