/*
 * @Author: Liang zh.liang.cn@gmail.com
 * @Date: 2022-08-27 21:23:51
 * @LastEditors: Liang zh.liang.cn@gmail.com
 * @LastEditTime: 2022-09-04 23:13:38
 * @Description:
 */
package longest_inc_seq

import "testing"

func TestLengthOfLIS(t *testing.T) {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}

	len := LengthOfLIS(arr)

	if len != 4 {
		t.Errorf("LengthOfLIS(%v) should be 4, got %v", arr, len)
	}
}

func TestLengthOfLIS_V2(t *testing.T) {
	arr := []int{10, 9, 2, 5, 3, 7, 101, 18}

	len := LengthOfLIS_V2(arr)

	if len != 4 {
		t.Errorf("LengthOfLIS_V2(%v) should be 4, got %v", arr, len)
	}
}
