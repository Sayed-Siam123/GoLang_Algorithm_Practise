package main

import (
	"fmt"
)

func main() {
	fmt.Println("sasasasasassasasasasasas")

	var a []int

	a = append(a, 3, 7, 8, 9, 12, 14, 31)

	sliceExample(a, 4)
}

func sliceExample(numberList []int, n int) {
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 512)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 1024)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))

	fmt.Println("\n")

	numberList = append(numberList, 20121)
	fmt.Println(numberList)
	fmt.Println(len(numberList))
	fmt.Println(cap(numberList))
}

func printLine(a int) {
	fmt.Println(a)
}

//CAPACITY is underlying array length, it will re-size based on slice length//
//LENGTH is how much element is inside the array list//
