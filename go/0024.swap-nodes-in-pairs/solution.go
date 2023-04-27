// Created by Vintingb at 2023/04/27 03:24
// https://leetcode.cn/problems/swap-nodes-in-pairs/

/*
24. 两两交换链表中的节点 (Medium)
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。

**示例 1：**

![](https://assets.leetcode.com/uploads/2020/10/03/swap_ex1.jpg)

```
输入：head = [1,2,3,4]
输出：[2,1,4,3]

```

**示例 2：**

```
输入：head = []
输出：[]

```

**示例 3：**

```
输入：head = [1]
输出：[1]

```

**提示：**

- 链表中节点的数目在范围 `[0, 100]` 内
- `0 <= Node.val <= 100`
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=begin

func swapPairs(head *ListNode) *ListNode {
	p := new(ListNode)
	cur := p
	for head != nil {
		if head.Next == nil {
			cur.Next = head
			return p.Next
		}
		one, two := head, head.Next
		head = head.Next.Next
		two.Next = one
		one.Next = nil // 这里必须将第一个节点的Next置空，否则偶数个节点情况下的最后两个节点会生成环
		cur.Next = two
		cur = one
	}
	return p.Next
}

// @lc code=end
