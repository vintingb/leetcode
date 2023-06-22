// Created by Vintingb at 2023/06/21 22:07
// https://leetcode.cn/problems/P5rCT8/

/*
剑指 Offer II 053. 二叉搜索树中的中序后继 (Medium)
给定一棵二叉搜索树和其中的一个节点 `p` ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 `null` 。

节点 `p` 的后继是值比 `p.val` 大的节点中键值最小的节点，即按中序遍历的顺序节点 `p` 的下一个节点。

**示例 1：**

![](https://assets.leetcode.com/uploads/2019/01/23/285_example_1.PNG)

```
输入：root = [2,1,3], p = 1
输出：2
解释：这里 1 的中序后继是 2。请注意 p 和返回值都应是 TreeNode 类型。

```

**示例 2：**

![](https://assets.leetcode.com/uploads/2019/01/23/285_example_2.PNG)

```
输入：root = [5,3,6,2,4,null,null,1], p = 6
输出：null
解释：因为给出的节点没有中序后继，所以答案就返回 null 了。

```

**提示：**

- 树中节点的数目在范围 `[1, 10⁴]` 内。
- `-10⁵ <= Node.val <= 10⁵`
- 树中各节点的值均保证唯一。

注意：本题与主站 285 题相同：
[https://leetcode-cn.com/problems/inorder-successor-in-bst/](https://leetcode-cn.com/problems/inorder-successor-in-bst/)
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=begin

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	var res *TreeNode
	for root != nil {
		if root.Val > p.Val {
			res = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return res
}

// @lc code=end
