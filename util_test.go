package BplusTree

import (
	"testing"
)

func makeInterfaceList(data interface{}) []interface{} {
	var res []interface{}
	switch data.(type) {
	case []int:
		for i := 0; i < len(data.([]int)); i++ {
			res = append(res, data.([]int)[i])
		}
	}
	return res
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		keys []interface{}
		key  interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "left edge", args: args{keys: makeInterfaceList([]int{1, 2, 3, 4, 5, 6, 7}), key: 0}, want: -1},
		{name: "right edge", args: args{keys: makeInterfaceList([]int{1, 2, 3, 4, 5, 6, 7}), key: 8}, want: -8},
		{name: "middle &not found", args: args{keys: makeInterfaceList([]int{1, 2, 3, 8, 9, 10, 11}), key: 4}, want: -4},
		{name: "find", args: args{keys: makeInterfaceList([]int{1, 2, 3, 8, 9, 10, 11}), key: 3}, want: 2},
	}

	BinarySearch := generateKeyBinarySearchFunc(nil, 0)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.keys, tt.args.key, 7); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
