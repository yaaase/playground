package main

import "fmt"

type reducer func(int, int) int
type filterer func(int) bool

func reduceInt(f reducer, arr []int) int {
	var next, result int
	result = arr[0]
	for i := 1; i < len(arr); i++ {
		next = arr[i]
		result = f(result, next)
	}
	return result
}

func filterInt(f filterer, arr []int) []int {
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
