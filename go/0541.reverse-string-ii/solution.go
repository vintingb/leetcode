// Created by Vintingb at 2023/05/01 05:22
// https://leetcode.cn/problems/reverse-string-ii/

/*
541. 反转字符串 II (Easy)
给定一个字符串 `s` 和一个整数 `k`，从字符串开头算起，每计数至 `2k` 个字符，就反转这 `2k` 字符中的前
`k` 个字符。

- 如果剩余字符少于 `k` 个，则将剩余字符全部反转。
- 如果剩余字符小于 `2k` 但大于或等于 `k` 个，则反转前 `k` 个字符，其余字符保持原样。

**示例 1：**

```
输入：s = "abcdefg", k = 2
输出："bacdfeg"

```

**示例 2：**

```
输入：s = "abcd", k = 2
输出："bacd"

```

**提示：**

- `1 <= s.length <= 10⁴`
- `s` 仅由小写英文组成
- `1 <= k <= 10⁴`
*/

package main

// @lc code=begin

func reverseStr(s string, k int) string {
	b := []byte(s)
	l := len(b)
	index := 0
	for index < l {
		if l-index >= k {
			for i, j := index, index+k-1; i < j; {
				b[i], b[j] = b[j], b[i]
				i++
				j--
			}
		} else {
			for i, j := index, l-1; i < j; {
				b[i], b[j] = b[j], b[i]
				i++
				j--
			}
		}
		index += 2 * k
	}
	return string(b)
}

// @lc code=end
