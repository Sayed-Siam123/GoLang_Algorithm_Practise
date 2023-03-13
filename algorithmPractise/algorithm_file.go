package main

import (
	"fmt"
	"sort"
)

func main() {
	var binarySearchVal []int
	var selectionSortVal []int
	var bubbleSortList []int

	var arrayForSum []int32

	var alice []int32
	var bob []int32

	var arrayForBigSum []int64

	var birthdayCakeUnitList []int32

	binarySearchVal = append(binarySearchVal, 3, 7, 8, 9, 12, 14, 31)
	selectionSortVal = append(selectionSortVal, 2, 8, 10, 5, 7)
	bubbleSortList = append(bubbleSortList, 3, 7, 31, 14, 12, 8, 9)
	arrayForSum = append(arrayForSum, 3, 4, 5, 6)

	alice = append(alice, 17, 28, 30)
	bob = append(bob, 99, 16, 8)

	arrayForBigSum = append(arrayForBigSum, 1000000001, 1000000002, 1000000003, 1000000004, 1000000005)

	birthdayCakeUnitList = append(birthdayCakeUnitList, 6, 6, 5, 6)

	//binarySearch(binarySearchVal, len(binarySearchVal), 30)
	//selectionSort(selectionSortVal, len(selectionSortVal))
	//bubbleSort(bubbleSortList, len(bubbleSortList))

	//hacker rank problem
	//simpleArraySum(arrayForSum)
	//compareTriplets(alice, bob)
	//aVeryBigSum(arrayForBigSum)
	birthdayCakeCandles(birthdayCakeUnitList)
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

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var result int32
	result = 0

	for i := 0; i < len(ar); i++ {
		result = result + ar[i]
	}

	fmt.Println(int(result))

	return result
}

func compareTriplets(a []int32, b []int32) []int32 {
	var alicePoints, bobPoints int32
	alicePoints, bobPoints = 0, 0
	var result []int32

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alicePoints++
		}
		if a[i] < b[i] {
			bobPoints++
		}
	}
	result = append(result, alicePoints, bobPoints)
	return result
}

func aVeryBigSum(ar []int64) int64 {
	var result int64
	result = 0

	for i := 0; i < len(ar); i++ {
		result = result + ar[i]
	}

	fmt.Println(int(result))

	return result

}

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	fmt.Println(candles)
	n := len(candles)
	count := 1
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n-i-1; j++ {
	//		if candles[j] < candles[j+1] {
	//			temp := candles[j]
	//			candles[j] = candles[j+1]
	//			candles[j+1] = temp
	//		}
	//	}
	//}

	//for i := 0; i < n-1; i++ {
	//	minimumIndex := i
	//	for j := i + 1; j < n; j++ {
	//		if candles[j] > candles[minimumIndex] {
	//			minimumIndex = j
	//		}
	//	}
	//
	//	if minimumIndex != i {
	//		temp := candles[i]
	//		candles[i] = candles[minimumIndex]
	//		candles[minimumIndex] = temp
	//	}
	//
	//}

	sort.SliceStable(candles, func(i, j int) bool {
		//i,j are represented for two value of the slice .
		return candles[i] > candles[j]
	})

	for i := 1; i < n; i++ {
		if candles[i] == candles[0] {
			count++
		}
	}

	fmt.Println(count)
	return int32(count)
}
