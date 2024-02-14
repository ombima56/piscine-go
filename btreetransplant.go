package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if root == node {
		return rplc
	}
	
	parent := BTreeSearchParent(root, node)
	if parent.Left == node {
		parent.Left = rplc
	}
	if parent.Right == node {
		parent.Right = rplc
	}
	return root
}
