package BplusTree

import (
	"strconv"
	"testing"
)

func Test_bPlusTree_Insert(t *testing.T) {
	order := 10
	type Goods struct {
		price int
		name  string
	}
	tree := InitBPlusTree(order, nil, Goods{}.price)
	for i := 0; i < order; i += 2 {
		tree.Insert(i, &Goods{i, "商品 " + strconv.Itoa(i)})
	}
	for i := 1; i < order; i += 2 {
		tree.Insert(i, &Goods{i, "商品 " + strconv.Itoa(i)})
	}
	tests := []struct {
		key  int
		want string
	}{
		{0, "商品 0"},
		{1, "商品 1"},
		{2, "商品 2"},
		{3, "商品 3"},
		{4, "商品 4"},
		{5, "商品 5"},
		{6, "商品 6"},
		{7, "商品 7"},
		{8, "商品 8"},
		{9, "商品 9"},
	}
	for _, tt := range tests {
		if name := tree.Search(tt.key).(*Goods).name; name != tt.want {
			t.Errorf("Search() = %v, want %v", name, tt.want)
		}
	}
}
