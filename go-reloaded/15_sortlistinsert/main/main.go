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
	var link2 *student.NodeI	// for tests

	// link = listPushBack(link, 1)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 9)

	PrintList(link)

	link = student.SortListInsert(link, -2)
	link = student.SortListInsert(link, 2)
	PrintList(link)

	fmt.Println()

	// tests

	link2 = listPushBack(link2, 0)
	link2 = listPushBack(link2, 12)
	link2 = listPushBack(link2, 20)
	link2 = listPushBack(link2, 23)
	link2 = listPushBack(link2, 23)
	link2 = listPushBack(link2, 24)
	link2 = listPushBack(link2, 30)
	link2 = listPushBack(link2, 41)
	link2 = listPushBack(link2, 53)
	link2 = listPushBack(link2, 57)
	link2 = listPushBack(link2, 59)
	
	PrintList(link2)

	link2 = student.SortListInsert(link2, 38)
	PrintList(link2)
}
