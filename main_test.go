package main

import (
	"reflect"
	"testing"
)

func TestReduceInts(t *testing.T) {
	type args struct {
		f   func(int, int) int
		arr []int
	}

	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"reduce add [] is 0", args{add, []int{}}, 0},
		{"reduce add [1] is 6", args{add, []int{1}}, 1},
		{"reduce add [1,2,3] is 6", args{add, []int{1, 2, 3}}, 6},
		{"reduce multiply [2,3,4] is 24", args{multiply, []int{2, 3, 4}}, 24},
		{"reduce max [4,1,9,4,11,4,11,3 is 11", args{max, []int{4, 1, 9, 4, 11, 4, 11, 3}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduceInt(tt.args.f, tt.args.arr); got != tt.want {
				t.Errorf("ReduceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInts(t *testing.T) {
	type args struct {
		f   func(int) bool
		arr []int
	}
	even := func(x int) bool {
		if x%2 == 0 {
			return true
		}
		return false
	}
	odd := func(x int) bool {
		return !even(x)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"filter even [1,2,3,4] is [2,4]", args{even, []int{1, 2, 3, 4}}, []int{2, 4}},
		{"filter odd [1,2,3,4] is [3,4]", args{odd, []int{1, 2, 3, 4}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterInt(tt.args.f, tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
