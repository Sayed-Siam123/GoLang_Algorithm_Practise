package main

import (
	"fmt"
)

func main() {

	val := 10
	val1 := 10
	val2 := 20

	doubleValue(val)
	fmt.Printf("The number is %d\n", val)
	doubleValuePtr(&val)
	fmt.Printf("The number is %d\n", val)

	add, sub := addAndSub(val2, val1)

	fmt.Printf("The add is %d and sub is %d\n", add, sub)

	result, err := errorReturn(-1, val1)

	if err != nil {
		fmt.Printf("ERROR %v\n", err)
	} else {
		fmt.Printf("The result is %d\n", result)
	}

}

func doubleValue(a int) {
	a *= 2
}

func doubleValuePtr(a *int) {
	*a *= 2
}

func addAndSub(a int, b int) (int, int) {
	return a + b, a - b
}

func errorReturn(a int, b int) (int, error) {
	//return a + b, fmt.Errorf("%v", error)
	if a < 0 {
		return 0, fmt.Errorf("%v", a)
	}

	return a + b, nil
}

//if I have to use exact value from memory I need to pass the value as pointer
// I can use pointer, or I can use return to return the values for same memory
