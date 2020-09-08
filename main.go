package main

import (
	"fmt"
	"lileetcode.com/lettnode"
)

type V struct {
	i int32
	j int64
	k int64
}

type Peak struct {
	Name      string
	Elevation int // in feet
}

func main() {
	l1 := lettnode.ListNode{
		Val: 1,
		Next: &lettnode.ListNode{
			Val: 2,
			Next: &lettnode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := lettnode.ListNode{
		Val: 1,
		Next: &lettnode.ListNode{
			Val: 3,
			Next: &lettnode.ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	lists := lettnode.MergeTwoListsRecursion(&l1, &l2)
	for lists != nil {
		fmt.Printf("%d-", lists.Val)
		lists = lists.Next
	}
	//a:= []int{1,2}
	//a = a[1:]
	//fmt.Println(a)
	//valid := match.IsValidUpgrade("[](])")
	//fmt.Println(valid)
	//peaks := []Peak{
	//	{"Aconcagua", 22838},
	//	{"Denali", 20322},
	//	{"Kilimanjaro", 19341},
	//	{"Mount Elbrus", 18510},
	//	{"Mount Everest", 29029},
	//	{"Mount Kosciuszko", 7310},
	//	{"Mount Vinson", 16050},
	//	{"Puncak Jaya", 16024},
	//}
	//
	//// does an in-place sort on the peaks slice, with tallest peak first
	//sort.Slice(peaks, func(i, j int) bool {
	//	return peaks[i].Elevation > peaks[j].Elevation
	//})
	//fmt.Println(peaks)
}
