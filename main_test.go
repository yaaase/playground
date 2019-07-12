package main

import "testing"

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
		{"reduce add [1,2,3] is 6", args{add, []int{1, 2, 3}}, 6},
		{"reduce multiply [2,3,4] is 24", args{multiply, []int{2, 3, 4}}, 24},
		{"reduce max [4,1,9,4,11,4,11,3 is 11", args{max, []int{4, 1, 9, 4, 11, 4, 11, 3}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceInts(tt.args.f, tt.args.arr); got != tt.want {
				t.Errorf("ReduceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
