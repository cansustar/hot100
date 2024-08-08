package normalarray

/*
轮转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
eg:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
*/

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

func RotateArray(nums []int, k int) []int {
	/*
		1. 反转整个数组
		2. 反转前k个元素
		3. 反转后n-k个元素
		时间复杂度O(n), 空间复杂度O(1)
	*/
	n := len(nums)
	k %= n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
	return nums
}
