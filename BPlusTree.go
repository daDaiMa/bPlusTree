package BplusTree

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
	nodeType       NodeType
	parentKeyIndex int
	parent         *treeNonLeafNode
	linkHead       *link
	size           int //保存当前key的size
}

type treeNode struct {
	nodeComm
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
			LEAF_NODE,
			-1,
			nil,
			newLink(),
			0},
		make([]interface{}, order),
		make([]interface{}, order),}
}
