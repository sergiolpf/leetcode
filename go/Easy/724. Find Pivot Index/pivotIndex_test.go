package main

import "testing"

func Test_pivotIndex(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{"teste 1", args{[]int{1, 7, 3, 6, 5, 6}}, 3},
		//{"teste 2", args{[]int{1, 2, 3}}, -1},
		//{"teste 3", args{[]int{2, 1, -1}}, 0},
		{"teste 4", args{[]int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex2(tt.args.numbers); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
