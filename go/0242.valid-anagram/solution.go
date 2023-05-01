// Created by Vintingb at 2023/04/28 00:39
// https://leetcode.cn/problems/valid-anagram/

/*
242. 有效的字母异位词 (Easy)
给定两个字符串 `s` 和 `t` ，编写一个函数来判断 `t` 是否是 `s` 的字母异位词。

**注意：** 若 `s` 和 `t` 中每个字符出现的次数都相同，则称 `s` 和 `t` 互为字母异位词。

**示例 1:**

```
输入: s = "anagram", t = "nagaram"
输出: true

```

**示例 2:**

```
输入: s = "rat", t = "car"
输出: false
```

**提示:**

- `1 <= s.length, t.length <= 5 * 10⁴`
- `s` 和 `t` 仅包含小写字母

**进阶:** 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

package main

// @lc code=begin

func isAnagram(s string, t string) bool {
	m := make([]int, 26)
	for _, r := range s {
		m[r-'a']++
	}
	for _, r := range t {
		if m[r-'a'] <= 0 {
			return false
		}
		m[r-'a']--
	}
	for _, r := range m {
		if r != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
