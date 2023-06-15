package main

import (
	"fmt"
)

func main() {
	var l1 = []int{1, 4, 7}
	l2 := []int{11, 12, 1, 1}
	a := mergeTwoLists(l1, l2)
	fmt.Println(a)

}
func mergeTwoLists(list1 []int, list2 []int) []int {

	if list1 == nil && list2 != nil {
		return list2
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	if list1 == nil && list2 != nil {
		return nil
	}
	for i := 0; i < len(list1); i++ {

	}

}
