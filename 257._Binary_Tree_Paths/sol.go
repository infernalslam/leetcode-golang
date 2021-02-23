/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	res := binaryTreeNums(root)
	return res
}

func binaryTreeNums(root *TreeNode) []string {

	if root == nil {
		return []string{}
	}

	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	res := []string{}

	tempLeft := binaryTreeNums(root.Left)

	for i := 0; i < len(tempLeft); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tempLeft[i])
	}

	tempRight := binaryTreeNums(root.Right)

	for i := 0; i < len(tempRight); i++ {
		res = append(res, strconv.Itoa(root.Val)+"->"+tempRight[i])
	}

	return res

}