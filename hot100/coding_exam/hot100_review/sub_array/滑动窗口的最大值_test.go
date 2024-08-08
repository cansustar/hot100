package subarray

import "testing"

/*
滑动窗口的最大值
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字(其实就是窗口大小为k)。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
eg:
	输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
	输出：[3,3,5,5,6,7]
*/

func SlidingWindowMaxNum(nums []int, k int) []int {
	/*
		进阶解法，使用双端队列，时间复杂度为O(n)
	*/
	n := len(nums)
	if n*k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	// 初始化双端队列。 存放的是元素的索引，而不是元素本身。 并且队列中的元素是递减的
	dequeue := make([]int, 0)
	result := make([]int, 0)
	// i就是当前遍历的元素的索引，也就是窗口结束的位置，即右边界
	for i, num := range nums {
		// 如果右边界位置的元素大于双端队列的最后一个元素，那么就将双端队列的最后一个元素出队，直到双端队列为空或者最后一个元素大于当前元素
		for len(dequeue) > 0 && num > nums[dequeue[len(dequeue)-1]] {
			// 将最后一个元素出队
			dequeue = dequeue[:len(dequeue)-1]
		}
		// 将当前元素的索引入队
		dequeue = append(dequeue, i)
		// 处理左边界，如果左边界的索引(i-k)等于双端队列的第一个元素的索引，那么就将第一个元素出队
		if i >= k && dequeue[0] == i-k {
			dequeue = dequeue[1:]
		}
		// 当i>=k-1时，也就是窗口大小为k时，将双端队列的第一个元素加入到结果中
		if i >= k-1 {
			result = append(result, nums[dequeue[0]])
		}
	}
	return result
}

func slidingWindowMaxNum(nums []int, k int) []int {
	/*	普通解法，需要遍历每个窗口，时间复杂度为O(n*k)
		套用滑动窗口的模板，只需要在窗口大小为k时，处理窗口内的元素，求最大值即可。
	*/
	n := len(nums)
	if n*k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	left, right := 0, 0
	window := make(map[int]int)
	var result []int
	for right < n {
		window[nums[right]]++
		right++
		if right-left == k {
			result = append(result, getMaxValue(window))
			delete(window, nums[left])
			left++
		}
	}
	return result
}

func getMaxValue(windows map[int]int) int {
	maxValue := 0
	for k := range windows {
		if k > maxValue {
			maxValue = k
		}
	}
	return maxValue
}

func TestName(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := SlidingWindowMaxNum(nums, k)
	t.Log(result)
}
