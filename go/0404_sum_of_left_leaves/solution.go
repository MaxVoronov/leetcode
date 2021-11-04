package sum_of_left_leaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return adder(root, 0)
}

func adder(node *TreeNode, sum int) int {
	if node != nil {
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		sum = adder(node.Left, sum)
		sum = adder(node.Right, sum)
	}

	return sum
}
