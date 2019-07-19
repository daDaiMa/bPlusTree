package BplusTree

import (
	"math/rand"
	"strconv"
	"testing"
)

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
