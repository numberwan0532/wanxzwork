package main

import (
	"sort"
)

// 合并区间
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	// 先对区间按照起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{}
	for _, interval := range intervals {
		// 如果merged为空，或者当前区间的起始位置大于merged中最后一个区间的结束位置，说明没有重叠，直接添加
		if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
			merged = append(merged, interval)
		} else {
			// 否则，存在重叠，更新merged中最后一个区间的结束位置
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
