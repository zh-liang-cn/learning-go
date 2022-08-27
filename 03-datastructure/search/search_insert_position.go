package search

/**
 * https://leetcode.cn/problems/search-insert-position/
 */

func searchInsert(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func searchInsertRecursion(nums []int, lo int, hi int, target int) int {
	mid := lo + (hi-lo)>>1

	if lo > hi {
		return lo
	}

	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		return searchInsertRecursion(nums, lo, mid-1, target)
	} else {
		return searchInsertRecursion(nums, mid+1, hi, target)
	}
}
