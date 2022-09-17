/*
 * @Author: Liang zh.liang.cn@gmail.com
 * @Date: 2022-08-28 01:07:56
 * @LastEditors: Liang zh.liang.cn@gmail.com
 * @LastEditTime: 2022-09-18 00:18:09
 * @Description: 最大和子序列问题，Max Sum Sequence，https://leetcode.cn/problems/maximum-subarray/
 */
package max_sum_seq

func MaxSumSeq1(nums []int) int {
	max := nums[0]

	for i := 0; i < len(nums); i++ { //{-1, -2},
		tmpMax, sum := nums[i], nums[i]

		for j := i + 1; j < len(nums); j++ {
			sum = sum + nums[j]
			if sum > tmpMax {
				tmpMax = sum
			}
		}

		if tmpMax > max {
			max = tmpMax
		}

	}

	return max
}

func MaxSumSeq2(nums []int) int {
	len := len(nums)
	dp := make([]int, len)

	dp[0] = nums[0]
	max := dp[0]

	for i := 1; i < len; i++ {
		tmp1 := dp[i-1] + nums[i]
		tmp2 := nums[i]
		if tmp1 > tmp2 {
			dp[i] = tmp1
		} else {
			dp[i] = tmp2
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
