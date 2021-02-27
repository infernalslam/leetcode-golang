/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := sol(root, low, high)
	return ans
}

func sol(root *TreeNode, l, h int) int {

	sum := 0

	if l <= root.Val && root.Val <= h {
		sum += root.Val
	}

	if root.Left != nil {
		sum += sol(root.Left, l, h)
	}

	if root.Right != nil {
		sum += sol(root.Right, l, h)
	}

	return sum
}