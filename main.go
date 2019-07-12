package main

import "fmt"

// ReduceInts ...
func ReduceInts(f func(int, int) int, arr []int) int {
	var next, result int
	result = arr[0]
	for i := 1; i < len(arr); i++ {
		next = arr[i]
		result = f(result, next)
	}
	return result
}

func main() {
	fmt.Println("I am main")
}
