package main

import (
	"fmt"
)

func main() {
	var binarySearchVal []int
	var selectionSortVal []int

	binarySearchVal = append(binarySearchVal, 3, 7, 8, 9, 12, 14, 31)
	selectionSortVal = append(selectionSortVal, 2, 8, 10, 5, 7)

	//binarySearch(binarySearchVal, len(binarySearchVal), 30)
	selectionSort(selectionSortVal, len(selectionSortVal))
}

func selectionSort(numberList []int, n int) {
	fmt.Println(numberList)
	//fmt.Println(n)

	for i := 0; i < n-1; i++ {
		minimumIndex := i
		for j := i + 1; j < n; j++ {
			if numberList[j] < numberList[minimumIndex] {
				minimumIndex = j
			}
		}

		if minimumIndex != i {
			temp := numberList[i]
			numberList[i] = numberList[minimumIndex]
			numberList[minimumIndex] = temp
		}

	}

	fmt.Println(numberList)

}

func binarySearch(numberList []int, n int, x int) {

	numberFound := false
	left := 0
	right := n - 1

	for left <= right {

		mid := (left + right) / 2

		if numberList[mid] == x {
			numberFound = true
			break
		} else if numberList[mid] < x {
			left = mid + 1
			continue
		} else if numberList[mid] > x {
			right = mid - 1
			continue
		}
	}

	if numberFound {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}

}
