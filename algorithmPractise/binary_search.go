package main

import (
	"fmt"
)

func main() {
	var a []int

	a = append(a, 3, 7, 8, 9, 12, 14, 31)

	binarySearch(a, len(a), 31)
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
