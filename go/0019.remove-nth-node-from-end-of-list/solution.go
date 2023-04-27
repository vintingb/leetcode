// Created by Vintingb at 2023/04/27 04:41
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

/*
19. 删除链表的倒数第 N 个结点 (Medium)
给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

```

**示例 2：**

```
输入：head = [1], n = 1
输出：[]

```

**示例 3：**

```
输入：head = [1,2], n = 1
输出：[1]

```

**提示：**

- 链表中结点的数目为 `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

**进阶：** 你能尝试使用一趟扫描实现吗？
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=begin

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	f := head // fast
	s := head // slow
	for i := 0; i < n; i++ {
		f = f.Next
		// if n == len(head)
		if f == nil {
			return head.Next
		}
	}
	// 找到需要删除的元素位置的前一个位置
	for f.Next != nil {
		s = s.Next
		f = f.Next
	}
	s.Next = s.Next.Next
	return head
}

// @lc code=end
