/*
 * @Author: Liang zh.liang.cn@gmail.com
 * @Date: 2022-08-27 21:20:58
 * @LastEditors: Liang zh.liang.cn@gmail.com
 * @LastEditTime: 2022-09-04 23:19:42
 * @Description: 最长递增子序列问题，Longest Increasing Subsequence，简称LIS，https://leetcode.cn/problems/longest-increasing-subsequence/
 *
 */
package longest_inc_seq

import (
	"github.com/zh-liang-cn/learning-go/03-datastructure/dp"
)

func l(arr []int, i int) int {
	if i >= len(arr)-1 {
		return 1
	}

	max_len := 0
	for j := i + 1; j < len(arr); j++ {
		if arr[j] > arr[i] {
			max_len = dp.Max(l(arr, j)+1, max_len)
		}
	}

	return max_len
}

func LengthOfLIS(arr []int) int {
	max_len := 0
	for i := 0; i < len(arr); i++ {
		max_len = dp.Max(l(arr, i), max_len)
	}

	return max_len
}

/**
 * arr := {10, 9, 2, 5, 3, 7, 101, 18}
 */
func LengthOfLIS_V2(arr []int) (max int) {

	n := len(arr) // n = 8
	max = 0
	l := make([]int, n)

	for i := n - 1; i >= 0; i-- { // i = [7, 6, 5, 4, 3, 2, 1, 0]
		l[i] = 1
		for j := i + 1; j < n; j++ { // if i = 2, then j = [3, 4, 5, 6, 7]
			if arr[j] > arr[i] {
				if l[i] < l[j]+1 {
					l[i] = l[j] + 1
				}

			}
		}
		if l[i] > max {
			max = l[i]
		}
	}

	return max
}
