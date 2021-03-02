package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var sliceLength int
	var input int
	var unsorted = make([]int, 0, 3)

	fmt.Println("Specify length of the array: ")
	fmt.Scan(&sliceLength)

	for i := 0; i < sliceLength; i++ {
		fmt.Println("Enter element of the array: ")
		fmt.Scan(&input)
		unsorted = append(unsorted, input)
	}

	sliceSize := sliceLength / 4
	slice1 := unsorted[:sliceSize]
	slice2 := unsorted[sliceSize : 2*(sliceSize)]
	slice3 := unsorted[2*(sliceSize) : 3*(sliceSize)]
	slice4 := unsorted[3*(sliceSize):]

	fmt.Println("Partitioned arrays", slice1, slice2, slice3, slice4)

	var waitGroup sync.WaitGroup
	waitGroup.Add(4)
	go sortList(slice1)
	waitGroup.Done()
	go sortList(slice2)
	waitGroup.Done()
	go sortList(slice3)
	waitGroup.Done()
	go sortList(slice4)
	waitGroup.Done()
	waitGroup.Wait()

	newSlice := mergeAndSortList(slice1, slice2, slice3, slice4)
	fmt.Println("Ordered slice:", newSlice)
}

func sortList(list []int) []int {
	sort.Ints(list)
	return list
}

func mergeAndSortList(l1 []int, l2 []int, l3 []int, l4 []int) []int {
	slice := []int{}
	slice = append(l1, l2...)
	slice = append(slice, l3...)
	slice = append(slice, l4...)
	sort.Ints(slice)
	return slice
}
