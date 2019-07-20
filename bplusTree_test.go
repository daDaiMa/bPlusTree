package BplusTree

import (
	"fmt"
	"math/rand"
	"strconv"
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

func Test_bPlusTree_Insert(t *testing.T) {
	order := 5
	type Goods struct {
		price int
		name  string
	}
	m := make(map[int]bool)
	tree := InitBPlusTree(order, nil, Goods{}.price)
	for i := 0; i < 500; i++ {
		num := rand.Int() % 1000
		for {
			if _, ok := m[num]; ok {
				num = rand.Int() % 1000
			} else {
				m[num] = true
				break
			}
		}
		tree.Insert(num, &Goods{num, "test" + strconv.Itoa(num)})
	}
	for val, _ := range m {
		if price := tree.Search(val).(*Goods).price; price != val {
			t.Errorf("Search() = %v, want %v", price, val)
		}
	}
}

func Test_bPlusTree_Delete(t *testing.T) {
	order := 5
	type Goods struct {
		price int
		name  string
	}
	m := make(map[int]bool)
	tree := InitBPlusTree(order, nil, Goods{}.price)
	for i := 0; i < 10; i++ {
		num := rand.Int() % 1000
		for {
			if _, ok := m[num]; ok {
				num = rand.Int() % 1000
			} else {
				m[num] = true
				break
			}
		}
		tree.Insert(num, &Goods{num, "test" + strconv.Itoa(num)})
	}
	tree.PrintSimply()
	for val, _ := range m {
		fmt.Println("delete: " + strconv.Itoa(val))
		tree.Delete(val)
		tree.PrintSimply()
		fmt.Println("alfter delete ")
	}
}
