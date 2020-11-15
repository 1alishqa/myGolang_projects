package main

import (
	"fmt"

	student ".."
)

func PrintList(l *student.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *student.NodeI
	var link2 *student.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	// link2 = listPushBack(link2, -2)
	// link2 = listPushBack(link2, 9)

	PrintList(student.SortedListMerge(link2, link))

	fmt.Println()

	var link3 *student.NodeI
	var link4 *student.NodeI

	link3 = listPushBack(link3, 0)
	link3 = listPushBack(link3, 3)
	link3 = listPushBack(link3, 8)
	link3 = listPushBack(link3, 8)
	link3 = listPushBack(link3, 13)
	link3 = listPushBack(link3, 19)
	link3 = listPushBack(link3, 34)
	link3 = listPushBack(link3, 38)
	link3 = listPushBack(link3, 46)

	link4 = listPushBack(link4, 7)
	link4 = listPushBack(link4, 39)
	link4 = listPushBack(link4, 45)
	link4 = listPushBack(link4, 53)
	link4 = listPushBack(link4, 59)
	link4 = listPushBack(link4, 70)
	link4 = listPushBack(link4, 76)
	link4 = listPushBack(link4, 79)

	PrintList(student.SortedListMerge(link3, link4))

}
