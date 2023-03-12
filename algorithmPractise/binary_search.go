package main

import (
	"fmt"
)

func main() {
	var a []int
	var bubbleSortList []int

	a = append(a, 3, 7, 8, 9, 12, 14, 31)

	bubbleSortList = append(bubbleSortList, 3, 7, 31, 14, 12, 8, 9)

	//binarySearch(a, len(a), 31)

	bubbleSort(bubbleSortList, len(bubbleSortList))
}

func bubbleSort(A []int, n int) {
	fmt.Println(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if A[j] > A[j+1] {
				temp := A[j]
				A[j] = A[j+1]
				A[j+1] = temp
			}
		}
	}

	fmt.Println(A)

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
		} else {
			numberFound = false
			break
		}
	}

	if numberFound {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}

}
