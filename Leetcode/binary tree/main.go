package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
