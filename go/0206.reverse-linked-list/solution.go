// Created by Vintingb at 2023/04/27 01:39
// https://leetcode.cn/problems/reverse-linked-list/

/*
206. 反转链表 (Easy)
给你单链表的头节点 `head` ，请你反转链表，并返回反转后的链表。

**示例 1：**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

```

**示例 2：**

![](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```
输入：head = [1,2]
输出：[2,1]

```

**示例 3：**

```
输入：head = []
输出：[]

```

**提示：**

- 链表中节点的数目范围是 `[0, 5000]`
- `-5000 <= Node.val <= 5000`

**进阶：** 链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=begin

func reverseList(head *ListNode) *ListNode {
	p := new(ListNode)
	for head != nil {
		tmp := p.Next
		p.Next = head
		head = head.Next
		p.Next.Next = tmp
	}
	return p.Next
}

// @lc code=end
