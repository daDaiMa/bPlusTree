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

	leafSimpleInsert(node, key, value, insert)
	if node.size > tree.order {
		split := node.size / 2
		// 创建
		rightLeaf := newLeafNode(tree.order)
		// 拷贝
		copyLeafNode(node, rightLeaf, split)
		// link
		node.link.addNext(rightLeaf.link)
		// 绑定父节点
		tree.leafBindParent(node, rightLeaf)
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

func copyLeafNode(ori, target *treeLeafNode, split int) {
	j := 0
	for i := split; i < len(ori.keys); i++ {
		target.keys[j] = ori.keys[i]
		target.data[j] = ori.data[i]
		j++
	}
	target.size = j
	ori.size = split
}

func (tree *bPlusTree) leafBindParent(left, right *treeLeafNode) {
	if left.parent == nil && right.parent == nil {
		parent := newNonLeafNode(tree.order)
		parent.keys[0] = right.keys[0]
		parent.subPtr[0] = left
		parent.subPtr[1] = right
		parent.size++

		left.parent = parent
		left.parentIndex = -1
		right.parent = parent
		right.parentIndex = 0

		tree.root = parent

	} else if left.parent != nil {
		insert := left.parentIndex + 1
		tree.nonLeafNodeInsert(left.parent, right.keys[0], right, insert)
	}
}

func (tree *bPlusTree) nonLeafNodeInsert(parent *treeNonLeafNode, key, treeNode interface{}, insert int) {
	for i := parent.size; i >= insert; i-- {
		if i == insert {
			parent.keys[i] = key
			parent.subPtr[i+1] = treeNode
			parent.size++
			if leaf, ok := treeNode.(*treeLeafNode); ok {
				leaf.parent = parent
				leaf.parentIndex = insert
			} else {
				treeNode.(*treeNonLeafNode).parent = parent
				treeNode.(*treeNonLeafNode).parentIndex = insert
			}
			break

		} else {
			parent.keys[i] = parent.keys[i-1]
			parent.subPtr[i+1] = parent.subPtr[i]
		}
		if leaf, ok := parent.subPtr[i+1].(*treeLeafNode); ok {
			leaf.parentIndex++
		} else {
			if parent.subPtr[i+1] == nil {
				continue
			}
			parent.subPtr[i+1].(*treeNonLeafNode).parentIndex++
		}
	}
	if parent.size == tree.order {
		split := tree.order / 2
		right := newNonLeafNode(tree.order)
		copyNonLeafNode(parent, right, split)
		parent.link.addNext(right.link)
		tree.nonLeafNodeBindParent(parent, right)
	}
}

func copyNonLeafNode(left, right *treeNonLeafNode, split int) {
	left.size = split
	j := 0
	for i := split; i < len(left.keys); i++ {
		right.keys[j] = left.keys[i]
		right.subPtr[j+1] = left.subPtr[i+1]
		if leaf, ok := right.subPtr[j+1].(*treeLeafNode); ok {
			leaf.parent = right
			leaf.parentIndex = j
		} else {
			right.subPtr[j+1].(*treeNonLeafNode).parent = right
			right.subPtr[j+1].(*treeNonLeafNode).parentIndex = j
		}
		j++
	}
	right.size = j
}

func (tree *bPlusTree) nonLeafNodeBindParent(left, right *treeNonLeafNode) {
	if left.parent == nil && right.parent == nil {
		parent := newNonLeafNode(tree.order)
		parent.keys[0] = right.keys[0]
		parent.subPtr[0] = left
		parent.subPtr[1] = right
		parent.size++

		right.parent = parent
		left.parent = parent
		right.parentIndex = 0
		left.parentIndex = -1

		tree.root = parent

	} else if left.parent != nil {
		insert := left.parentIndex + 1
		tree.nonLeafNodeInsert(left.parent, right.keys[0], right, insert)
	}
}
