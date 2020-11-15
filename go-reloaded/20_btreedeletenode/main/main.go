package main

import (
	"fmt"
	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)

	fmt.Println()

	root2 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root2, "07")
	student.BTreeInsertData(root2, "05")
	student.BTreeInsertData(root2, "12")
	student.BTreeInsertData(root2, "10")
	student.BTreeInsertData(root2, "02")
	student.BTreeInsertData(root2, "03")
	node2 := student.BTreeSearchItem(root2, "02")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root2, fmt.Println)
	root2 = student.BTreeDeleteNode(root2, node2)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root2, fmt.Println)
}
