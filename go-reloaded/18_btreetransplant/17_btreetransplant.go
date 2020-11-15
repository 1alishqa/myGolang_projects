package student

type TreeNode struct {
	Left, Right, Parent * TreeNode
	Data string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}

func BTreeSearchItem(root *TreeNode, element string) *TreeNode {
	if root == nil {
		return nil
	}

	if element < root.Data {
		return BTreeSearchItem(root.Left, element)

	} else if element > root.Data {
		return BTreeSearchItem(root.Right, element)

	} else {
		return root
	}
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {

	if root != nil {

		BTreeApplyInorder(root.Left, f)
		f(root.Data)
		BTreeApplyInorder(root.Right, f)

	}
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	if root == nil {
		return &TreeNode{}
	}

	if root.Data > node.Data {
		if root.Left.Data == node.Data {
			root.Left = rplc
			rplc.Parent = root
			return root

		} else {
			BTreeTransplant(root.Left, node, rplc)
		}

	} else if root.Data < node.Data {
		if root.Right.Data == node.Data {
			root.Right = rplc
			rplc.Parent = root
			return root

		} else {
			BTreeTransplant(root.Right, node, rplc)
		}

	} else {
		return rplc
	}
	return root
	
}
