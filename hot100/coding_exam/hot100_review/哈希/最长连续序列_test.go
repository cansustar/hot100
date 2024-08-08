package 哈希

import "testing"

/*
	最长连续序列

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
eg:
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。
*/
func LongestConsecutiveSequence(nums []int) int {
	/*
		可以使用hash表来存储每个元素是否存在。
		如果当前数字的前一个数字不存在，说明当前数字是一个连续序列的起点。
		然后不断检查下一个元素是否在hash表中，直到找不到，尝试更新当前最大序列长度。
		时间复杂度O(n), 空间复杂度O(n)
	*/
	// 初始化连续序列最大长度
	strLen := 0
	// int表示元素值，bool表示元素是否存在
	hashTable := make(map[int]bool)
	// 将所有元素存入hash表
	for _, num := range nums {
		hashTable[num] = true
	}
	// 遍历hash表
	for num := range hashTable {
		// 判断当前元素的前一个元素是否存在，如果不存在，说明当前元素是一个连续序列的起点
		if !hashTable[num-1] {
			// 更新正在处理的元素
			curNum := num
			// 由于当前元素是连续序列的起点，所以当前长度为1
			curLen := 1
			// 不断检查下一个元素是否在hash表中，直到找不到
			for hashTable[curNum+1] {
				// 更新当前元素
				curNum++
				// 更新当前长度
				curLen++
			}
			// 当找不到下一个元素时，会跳出for循环，然后更新最大长度
			// 更新最大长度
			//str_len = max(str_len, cur_len)
			if curLen > strLen {
				strLen = curLen
			}
		}
	}
	return strLen
}

func TestLongestSequence(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	result := LongestConsecutiveSequence(nums)
	t.Log(result)
}
