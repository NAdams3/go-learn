package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(node *TreeNode) int {
	// [][]treenode for walking a corresponding sum [] for result
	walk := make([][]*TreeNode, 0)
	sums := make([]int, 0)

	walk = append(walk, []*TreeNode{node})

	max := 0
	idx := 0

	for i := 0; i < len(walk); i++ {
		levelNodes := walk[i]
		sum := 0
		for j, node := range levelNodes {
			sum += node.Val
			if j == 0 {
				if node.Left != nil {
					walk = append(walk, []*TreeNode{node.Left})
				}

				if node.Right != nil {
					walk = append(walk, []*TreeNode{node.Right})
				}
			} else {
				if node.Left != nil {
					walk[i+1] = append(walk[i+1], node.Left)
				}
				if node.Right != nil {
					walk[i+1] = append(walk[i+1], node.Right)
				}
			}

		}
		sums = append(sums, sum)
	}

	for k, sum := range sums {
		if sum > max {
			max = sum
			idx = k
		}
	}

	return idx

}
