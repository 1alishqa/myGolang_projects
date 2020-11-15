package main

import (
	"fmt"
	student ".."
)

func main() {
/*
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)*/
	
	
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "11")
	student.BTreeInsertData(root, "10")
	student.BTreeApplyInorder(root, fmt.Println)
	fmt.Println()
	node := student.BTreeSearchItem(root, "10")
	replacement := &student.TreeNode{Data: "8"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)

	fmt.Println()
	fmt.Println()

	/*
	root2 := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root2, "07")
	student.BTreeInsertData(root2, "12")
	student.BTreeInsertData(root2, "10")
	student.BTreeInsertData(root2, "05")
	student.BTreeInsertData(root2, "02")
	student.BTreeInsertData(root2, "03")
	node2 := student.BTreeSearchItem(root2, "12")
	replacement2 := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement2, "60")
	student.BTreeInsertData(replacement2, "33")
	student.BTreeInsertData(replacement2, "12")
	student.BTreeInsertData(replacement2, "15")
	root2 = student.BTreeTransplant(root2, node2, replacement2)
	student.BTreeApplyInorder(replacement2, fmt.Println)
	*/
}
