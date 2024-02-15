package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Left == nil {
		root = transplant(root, node, node.Right)
	} else if node.Right == nil {
		root = transplant(root, node, node.Left)
	} else {
		min := minimum(node.Right)
		if min.Parent != node {
			root = transplant(root, min, min.Right)
			min.Right = node.Right
			min.Right.Parent = min
		}
		root = transplant(root, node, min)
		min.Left = node.Left
		min.Left.Parent = min
	}
	return root
}

func transplant(root, u, v *TreeNode) *TreeNode {
	if u.Parent == nil {
		root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
	return root
}

func minimum(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
