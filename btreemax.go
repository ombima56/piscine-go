package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := root
	for node.Right != nil {
		node = node.Right
	}
	return node
}
