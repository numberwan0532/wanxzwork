package main

// 两数之和
func findTowNumTotal(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{nums[i], nums[j]}
			}
		}
	}
	return []int{}
}

func findTowNumTotal1(nums []int, target int) []int {
	marr := make(map[int]int)
	for i, v := range nums {
		if val, ok := marr[target-v]; ok {
			return []int{nums[val], nums[i]}
		}
		marr[v] = i
	}
	return []int{}
}
