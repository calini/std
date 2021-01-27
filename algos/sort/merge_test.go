package sort

import (
	"testing"
)

func TestInts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Empty",
			args: args{
				arr: []int{},
			},
		},
		{
			name: "Single",
			args: args{
				arr: []int{1},
			},
		},
		{
			name: "Simple",
			args: args{
				arr: []int{1, 3, -2, 4},
			},
		},
		{
			name: "Complex",
			args: args{
				arr: []int{1, 3, -2, 4, -1, 33, -8, 4, 1, 3, -2, 4},
			},
		},
		{
			name: "Sorted",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
		},
		{
			name: "Inverse",
			args: args{
				arr: []int{4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := MergeSort(tt.args.arr)
			t.Logf("Array that goes in: %+v, array that gets out %+v\n", tt.args.arr, arr)

			if !IsSortedInts(arr) {
				t.Fail()
			}
		})
	}
}