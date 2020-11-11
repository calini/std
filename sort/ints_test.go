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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ints(tt.args.arr)
			if !IsSortedInts(tt.args.arr) {
				t.Fail()
			}
		})
	}
}