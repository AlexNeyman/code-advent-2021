package util

import "sort"

func AbsInt(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}

func MaxInt(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func MinInt(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

func SumInts(nums []int) int {
	res := 0

	for _, num := range nums {
		res += num
	}

	return res
}

func GetBitValue(num int, shift int) int {
	return (num & (1 << shift)) >> shift
}
