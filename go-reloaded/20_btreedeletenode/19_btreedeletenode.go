package student

type TreeNode struct {
	Left, Right, Parent *TreeNode
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

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	if root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	if node.Data < root.Data {

		root.Left = BTreeDeleteNode(root.Left, node)

	} else if node.Data > root.Data {

		root.Right = BTreeDeleteNode(root.Right, node)

	} else {

		if root.Left == nil {
			
			temp := root.Right
			root = &TreeNode{}
			return temp

		} else if root.Right == nil {

			temp := root.Left
			root = &TreeNode{}
			return temp
		}

		temp := BTreeMin(root.Right)
		root.Data = temp.Data
		root.Right = BTreeDeleteNode(root.Right, BTreeMin(root.Right))

	}

	return root

}
