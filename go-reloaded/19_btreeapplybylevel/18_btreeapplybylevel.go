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

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))  {

	if root == nil {
		return
	}
	height := BTreeLevelCount(root)
	for i := 1; i <= height; i++ {
		Level(root, i, f)
	}
	
}

func Level(root *TreeNode, level int, f func(...interface{}) (int, error)) {

	if root == nil {
		return
	}

	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		Level(root.Left, level - 1, f)
		Level(root.Right, level - 1, f)
	}
	
}

func BTreeLevelCount(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := BTreeLevelCount(root.Left)
	right := BTreeLevelCount(root.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}

}
