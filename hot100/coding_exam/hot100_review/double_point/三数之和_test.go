package double_pointer

import (
	"sort"
	"testing"
)

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
返回所有和为 0 且不重复的三元组。
*/

func ThreeSum(nums []int) [][]int {
	// 要求三数和为0，可以转化为两数和为target的问题。
	// 遍历数组，将当前元素作为target，然后在剩下的元素中找两数和为-target的两个数。
	// 可以用双指针，如何让指针有目的地移动？ 即当两数和大于target时，移动某一指针使得和变小，反之亦然。
	// 所以需要先对数组进行排序
	// 时间复杂度：O(n^2)， 空间复杂度：O(1)
	sort.Ints(nums) // 排序后，数组元素就是从小到大排列的
	ans := make([][]int, 0)
	for i := range nums {
		// 如果当前元素已经>0，那么后面的元素肯定都>0，不可能存在三数和为0的情况
		if nums[i] > 0 {
			return ans
		}
		// 由于题目要求不可以重复，所以需要跳过重复的元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				// 由于题目要求答案中不可包含重复的三元组，所以需要跳过重复的元素
				// 当移动指针时，如果指针指向的元素与上一个元素相同，则需要跳过
				for left < right && nums[left] == nums[left+1] {
					// 继续移动
					left++
				}
				// 右指针同理
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSum(nums)
	t.Log(result)
}
