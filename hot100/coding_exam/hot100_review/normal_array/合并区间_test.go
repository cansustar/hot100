package normalarray

import (
	"math"
	"sort"
	"testing"
)

/*
 合并区间
 给出一个区间的集合，请合并所有重叠的区间。
eg:
输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
*/

func MergeInterval(intervals [][]int) [][]int {
	// 为了更方便地得到重叠的区间，我们首先将区间按照左端点排序
	// 然后遍历区间，如果当前区间的左端点大于前一个区间的右端点，说明两个区间不重叠，直接将当前区间加入结果集
	// 如果当前区间的左端点小于等于前一个区间的右端点，说明两个区间重叠，更新前一个区间的右端点
	// 时间复杂度O(nlogn), 空间复杂度O(n)

	// 初始化结果集
	ans := make([][]int, 0)
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 因为第一个区间肯定不会与前一个区间重叠，所以直接将第一个区间加入结果集
	ans = append(ans, intervals[0])

	// 遍历区间,因为第一个区间已经加入结果集，所以从第二个区间开始遍历
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的左端点，大于当前结果中，最后一个区间的右端点，则说明两个区间不重叠
		if intervals[i][0] > ans[len(ans)-1][1] {
			ans = append(ans, intervals[i])
		} else {
			// 如果当前区间的左端点小于等于当前结果中，最后一个区间的右端点，则说明两个区间重叠
			// 更新当前结果中，最后一个区间的右端点，为两个区间的右端点的最大值
			ans[len(ans)-1][1] = int(math.Max(float64(ans[len(ans)-1][1]), float64(intervals[i][1])))
		}
	}
	return ans
}

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := MergeInterval(intervals)
	t.Log(result)
}
