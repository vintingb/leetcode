// Created by Vintingb at 2023/04/28 01:20
// https://leetcode.cn/problems/intersection-of-two-arrays/

/*
349. 两个数组的交集 (Easy)
给定两个数组 `nums1` 和 `nums2` ，返回 它们的交集 。输出结果中的每个元素一定是 **唯一**
的。我们可以 **不考虑输出结果的顺序** 。

**示例 1：**

```
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

```

**示例 2：**

```
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的

```

**提示：**

- `1 <= nums1.length, nums2.length <= 1000`
- `0 <= nums1[i], nums2[i] <= 1000`
*/

package main

// @lc code=begin

func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0, len(nums1))
	m := make(map[int]struct{}, len(nums1))
	tmp := make(map[int]struct{}, len(nums1))
	for _, num := range nums1 {
		m[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := m[num]; ok {
			tmp[num] = struct{}{}
		}
	}
	for num := range tmp {
		ans = append(ans, num)
	}
	return ans
}

// @lc code=end
