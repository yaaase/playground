package main

import "testing"

func TestReduceInts(t *testing.T) {
	type args struct {
		f   func(int, int) int
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"reduce + [1,2,3] is 6", args{func(x, y int) int { return x + y }, []int{1, 2, 3}}, 6},
		{"reduce * [2,3,4] is 24", args{func(x, y int) int { return x * y }, []int{2, 3, 4}}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceInts(tt.args.f, tt.args.arr); got != tt.want {
				t.Errorf("ReduceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}
