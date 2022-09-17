/*
 * @Author: Liang zh.liang.cn@gmail.com
 * @Date: 2022-08-28 01:08:59
 * @LastEditors: Liang zh.liang.cn@gmail.com
 * @LastEditTime: 2022-09-17 21:20:24
 * @Description:
 */
package max_sum_seq

import (
	"testing"
)

var numsArray = [][]int{
	{-1},
	{-1, -2},
	{-2, 1, -3, 4, -1, 2, 1, -5, 4},
}
var resultsArray = []int{
	-1,
	-1,
	6,
}

func TestMaxSumSeq1(t *testing.T) {
	for i, e := range numsArray {
		result := MaxSumSeq1(e)
		if resultsArray[i] != result {
			t.Errorf("MaxSumSeq1(%v) = %v, want %d", e, result, resultsArray[i])
		}
	}
}

func TestMaxSumSeq2(t *testing.T) {
	for i, e := range numsArray {
		result := MaxSumSeq2(e)
		if resultsArray[i] != result {
			t.Errorf("MaxSumSeq2(%v) = %v, want %d", e, result, resultsArray[i])
		}
	}
}
