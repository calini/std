package inversions

import (
	"testing"
)

func TestMergeSortWithInversions(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty",
			args: args{
				arr: []int{},
			},
			want: 0,
		},
		{
			name: "No Inversions",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
			want: 0,
		},
		{
			name: "Only left and right inversions",
			args: args{
				arr: []int{2, 1, 3, 5, 4},
			},
			want: 2,
		},
		{
			name: "Only cross inversions",
			args: args{
				arr: []int{5, 4, 3, 2, 1},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := MergeSortWithInversions(tt.args.arr)
			if got != tt.want {
				t.Errorf("MergeSortWithInversions() got = %v, want %v", got, tt.want)
			}
		})
	}
}
