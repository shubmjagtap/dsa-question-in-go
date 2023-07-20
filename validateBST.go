/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"fmt"
	"math"
)

func isValidBST(root *TreeNode) bool {
	return isValidBSTUtil(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTUtil(root *TreeNode, minVal, maxVal int64) bool {
	if root == nil {
		return true
	}

	if int64(root.Val) <= minVal || int64(root.Val) >= maxVal {
		return false
	}

	return isValidBSTUtil(root.Left, minVal, int64(root.Val)) && isValidBSTUtil(root.Right, int64(root.Val), maxVal)
}
