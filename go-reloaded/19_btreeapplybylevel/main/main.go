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
	student.BTreeInsertData(root, "11")
	student.BTreeInsertData(root, "10")
	student.BTreeApplyByLevel(root, fmt.Println)

	fmt.Println()

	root2 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root2, "07")
	student.BTreeInsertData(root2, "05")
	student.BTreeInsertData(root2, "12")
	student.BTreeInsertData(root2, "02")
	student.BTreeInsertData(root2, "03")
	student.BTreeInsertData(root2, "10")
	student.BTreeApplyByLevel(root2, fmt.Println)

	fmt.Println()

	root3 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root3, "07")
	student.BTreeInsertData(root3, "05")
	student.BTreeInsertData(root3, "12")
	student.BTreeInsertData(root3, "02")
	student.BTreeInsertData(root3, "03")
	student.BTreeInsertData(root3, "10")
	student.BTreeApplyByLevel(root3, fmt.Print)

	fmt.Println()
}
