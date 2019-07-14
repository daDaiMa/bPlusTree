package BplusTree

func (tree *bPlusTree) Insert(key, value interface{}) {
	if tree.root == nil {
		leaf := newLeafNode(tree.order)
		leaf.keys[0] = key
		leaf.data[0] = value
		leaf.size++
		tree.root = leaf
	} else {
		node := tree.root
		for {
			if leaf, ok := node.(*treeLeafNode); ok {
				// 如果是叶子节点
				tree.leafInsert(leaf, key, value)
				return
			} else {
				//非叶子节点
				index := tree.binarySearch(node.(*treeNonLeafNode).keys, key, node.(*treeNonLeafNode).size)
				if index >= 0 {
					node = node.(*treeNonLeafNode).subPtr[index]
				} else {
					node = node.(*treeNonLeafNode).subPtr[-index-1]
				}
			}
		}
	}
}

func (tree *bPlusTree) leafInsert(node *treeLeafNode, key, value interface{}) {
	insert := tree.binarySearch(node.keys, key, node.size)
	if insert > 0 {
		// TODO:insert大于0 说明有key重复 后期要处理

	} else {
		insert = -insert - 1
	}
	if node.size+1 > tree.order {
		split := tree.order / 2
		//分成 左右两个节点 split之后全放在右边的节点里
		if insert < split {

		} else {

		}

	} else {
		leafSimpleInsert(node, key, value, insert)
	}
}

func leafSimpleInsert(node *treeLeafNode, key, value interface{}, insert int) {
	for i := node.size; i > insert; i-- {
		node.keys[i] = node.keys[i-1]
		node.data[i] = node.data[i-1]
	}
	node.keys[insert] = key
	node.data[insert] = value
	node.size++
}
