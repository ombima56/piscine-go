package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return BTreeIsBinaryUtil(root.Left, "", root.Data) && BTreeIsBinaryUtil(root.Right, root.Data, "")
}

func BTreeIsBinaryUtil(node *TreeNode, min, max string) bool {
	if node == nil {
		return true
	}
	return (min == "" || node.Data > min) && (max == "" || node.Data < max) &&
		BTreeIsBinaryUtil(node.Left, min, node.Data) && BTreeIsBinaryUtil(node.Right, node.Data, max)
}
