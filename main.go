package main

import "fmt"

type intReducer func(int, int) int
type intFilterer func(int) bool

func reduceInt(f intReducer, arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var next, result int
	result = arr[0]
	for i := 1; i < len(arr); i++ {
		next = arr[i]
		result = f(result, next)
	}
	return result
}

func filterInt(f intFilterer, arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	results := make([]int, 0)
	for _, i := range arr {
		if f(i) {
			results = append(results, i)
		}
	}
	return results
}

func main() {
	fmt.Println("I am main")
}
