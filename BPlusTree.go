package BplusTree

import (
	"unsafe"
)

type NodeType int

const (
	LEAF_NODE NodeType = iota
	NON_LEAF_NODE
)

const (
	MAX_ORDER = 100
)

type bPlusTree struct {
	order int         //阶数
	root  interface{} // 可能是 treeLeafNode 或是 treeNonLeafNode
	//compareFunc  func(a, b interface{}) bool
	binarySearch func(keys []interface{}, key interface{}, size int) int
}

type nodeComm struct {
	parent      *treeNonLeafNode
	parentIndex int
	link        *link
	size        int //保存当前key的size
}

type treeLeafNode struct {
	nodeComm
	keys []interface{}
	data []interface{}
}

type treeNonLeafNode struct {
	nodeComm
	keys   []interface{}
	subPtr []interface{} //仅两种可能 1.treeLeafNode 2.treeNonLeafNode
}

var leafFieldOffset uintptr
var nonLeafFieldOffset uintptr

func getLeaf(l **link) *treeLeafNode {
	if l == nil {
		return nil
	}
	return (*treeLeafNode)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) - leafFieldOffset))
}

func getNonLeaf(l **link) *treeNonLeafNode {
	if l == nil {
		return nil
	}
	return (*treeNonLeafNode)(unsafe.Pointer(uintptr(unsafe.Pointer(l)) - nonLeafFieldOffset))
}

func init() {
	dummy := &treeLeafNode{}
	leafFieldOffset = uintptr(unsafe.Pointer(&dummy.link)) - uintptr(unsafe.Pointer(dummy))
	nonLeafDummy := &treeNonLeafNode{}
	nonLeafFieldOffset = uintptr(unsafe.Pointer(&nonLeafDummy.link)) - uintptr(unsafe.Pointer(nonLeafDummy))
}

/*
 * 创建一颗指定阶数的B+树
 */
func InitBPlusTree(order int, compareFunc func(a, b interface{}) int, keyExample interface{}) *bPlusTree {
	if order < 3 || order > MAX_ORDER {
		panic("B+树的阶数不在范围内")
		return nil
	}
	binarySearch := generateKeyBinarySearchFunc(compareFunc, keyExample)
	return &bPlusTree{order, nil, binarySearch}
}

func newLeafNode(order int) *treeLeafNode {
	return &treeLeafNode{
		nodeComm{
			nil,
			-1,
			newLink(),
			0},
		make([]interface{}, order+1),
		make([]interface{}, order+1),}
	/*
	 * 为啥要make order+1个空间呢 ？
	 * 因为为了分裂方便
	 */
}

func newNonLeafNode(order int) *treeNonLeafNode {
	return &treeNonLeafNode{
		nodeComm{
			nil,
			-1,
			newLink(),
			0,
		},
		make([]interface{}, order),
		make([]interface{}, order+1),
	}
}
