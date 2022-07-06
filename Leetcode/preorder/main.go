package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	var result = []int{}
	if root != nil {
		Traverse(root, &result)
	}
	//sort result
	//sort.Ints(result)
	var sum = 0
	for i := 0; i < len(result); i++ {
		if result[i] >= low && result[i] <= high {
			sum = sum + result[i]
		}
	}
	return sum
}

func Traverse(current *TreeNode, result *[]int) {

	if current.Left != nil {
		Traverse(current.Left, result)
	}
	*result = append(*result, current.Val)
	if current.Right != nil {
		Traverse(current.Right, result)
	}

}
