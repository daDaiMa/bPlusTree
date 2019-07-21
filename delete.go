package BplusTree

func (tree *bPlusTree) Delete(key interface{}) {
	node := tree.root
	if node == nil {
		return
	}
	for {
		if leaf, ok := node.(*treeLeafNode); ok {
			tree.deleteFormLeaf(key, leaf)
			return
		} else {
			nonLeaf, _ := node.(*treeNonLeafNode)
			index := tree.binarySearch(nonLeaf.keys, key, nonLeaf.size)
			if index >= 0 {
				node = nonLeaf.subPtr[index+1]
			} else {
				node = nonLeaf.subPtr[-index-1]
			}
		}
	}
}

func (tree *bPlusTree) deleteFormLeaf(key interface{}, leaf *treeLeafNode) {
	index := tree.binarySearch(leaf.keys, key, leaf.size)
	if index < 0 {
		return
	}
	simpleDelete(leaf, index)
	if leaf.parent == nil {
		if leaf.size == 0 {
			tree.root = nil
		}
		return
	}
	if leaf.size < (tree.order+1)/2 {
		left := getLeaf(leaf.link.pre)
		right := getLeaf(leaf.link.next)
		if choiceLeft := makeSiblingChoice(left, right); choiceLeft {
			//选择左边的
			if left.size+leaf.size <= tree.order {
				//可以合并
				tree.mergeToLeftLeaf(left, leaf)
			} else {
				//从左边借一个
				shiftFromLeftLeaf(left, leaf)
			}
		} else {
			//选择右边
			if right.size+leaf.size <= tree.order {
				//合并
				tree.mergeToRightLeaf(leaf, right)
			} else {
				//从右边借一个
				shiftFromRightLeaf(leaf, right)
			}
		}
	}

}

func simpleDelete(leaf *treeLeafNode, index int) {
	for i := index + 1; i < leaf.size; i++ {
		leaf.keys[i-1] = leaf.keys[i]
		leaf.data[i-1] = leaf.data[i-1]
	}
	leaf.size--
	if index == 0 && leaf.parent != nil && leaf.parentIndex != -1 {
		replaceRecursive(leaf.parent, leaf.parentIndex, leaf.keys[0])
	}
}

func makeSiblingChoice(left, right interface{}) bool {
	switch left.(type) {
	case *treeLeafNode:
		if left.(*treeLeafNode) == nil {
			return false
		}
		if right.(*treeLeafNode) == nil {
			return true
		}
		if left.(*treeLeafNode).size > right.(*treeLeafNode).size {
			return true
		}
		return false
	case *treeNonLeafNode:
		if left.(*treeNonLeafNode) == nil {
			return false
		}
		if right.(*treeNonLeafNode) == nil {
			return true
		}
		if left.(*treeNonLeafNode).size > right.(*treeNonLeafNode).size {
			return true
		}
		return false
	default:
		return false
	}
}

func shiftFromLeftLeaf(left, leaf *treeLeafNode) {
	for i := leaf.size; i > 0; i-- {
		leaf.keys[i] = leaf.keys[i-1]
		leaf.data[i] = leaf.data[i-1]
	}
	leaf.keys[0] = left.keys[left.size-1]
	leaf.data[0] = left.data[left.size-1]
	leaf.size++
	left.size--
	replaceRecursive(leaf.parent, leaf.parentIndex, leaf.keys[0])
}

func shiftFromRightLeaf(leaf, right *treeLeafNode) {
	leaf.keys[leaf.size] = right.keys[0]
	leaf.data[leaf.size] = right.data[0]
	leaf.size++
	right.size--
	for i := 0; i < right.size; i++ {
		right.keys[i] = right.keys[i+1]
		right.data[i] = right.data[i+1]
	}
	replaceRecursive(right.parent, right.parentIndex, right.keys[0])
}
func replaceRecursive(nonLeaf *treeNonLeafNode, index int, key interface{}) {
	nonLeaf.keys[index] = key
	if index == 0 && nonLeaf.parent != nil && nonLeaf.parentIndex != -1 {
		replaceRecursive(nonLeaf.parent, nonLeaf.parentIndex, key)
	}
}
func (tree *bPlusTree) mergeToLeftLeaf(left, leaf *treeLeafNode) {
	for i := 0; i < leaf.size; i++ {
		left.keys[left.size] = leaf.keys[i]
		left.data[left.size] = leaf.data[i]
		left.size++
	}
	leaf.link.deleteSelf()
	tree.deleteFromNonLeaf(leaf.parent, leaf.parentIndex)
}
func (tree *bPlusTree) mergeToRightLeaf(leaf, right *treeLeafNode) {
	for i := 0; i < right.size; i++ {
		leaf.keys[leaf.size] = right.keys[i]
		leaf.data[leaf.size] = right.data[i]
		leaf.size++
	}
	right.link.deleteSelf()
	tree.deleteFromNonLeaf(right.parent, right.parentIndex)
}

func (tree *bPlusTree) deleteFromNonLeaf(nonLeaf *treeNonLeafNode, delete int) {
	simpleDeleteFromNonLeaf(nonLeaf, delete)
	if nonLeaf.parent == nil {
		if nonLeaf.size == 0 {
			tree.root = nonLeaf.subPtr[0]
			switch nonLeaf.subPtr[0].(type) {
			case *treeLeafNode:
				tree.root.(*treeLeafNode).parent = nil
			case *treeNonLeafNode:
				tree.root.(*treeNonLeafNode).parent = nil
			}
		}
		return
	}
	if nonLeaf.size < (tree.order-1)/2 {
		left := getNonLeaf(nonLeaf.link.pre)
		right := getNonLeaf(nonLeaf.link.next)
		if choiceLeft := makeSiblingChoice(left, right); choiceLeft {
			if left.size+nonLeaf.size <= tree.order-1 {
				tree.mergeLeftNonLeaf(left, nonLeaf)
			} else {
				shiftFromLeftNonLeaf(left, nonLeaf)
			}
		} else {
			if right.size+nonLeaf.size <= tree.order-1 {
				tree.mergeRightNonLeaf(nonLeaf, right)
			} else {
				shiftFromRightNonLeaf(nonLeaf, right)
			}
		}
	}
}

func (tree *bPlusTree) mergeLeftNonLeaf(left, nonLeaf *treeNonLeafNode) {
	for i := 0; i < nonLeaf.size; i++ {
		left.keys[left.size] = nonLeaf.keys[i]
		left.subPtr[left.size+1] = nonLeaf.subPtr[i+1]
		switch left.subPtr[left.size+1].(type) {
		case *treeNonLeafNode:
			left.subPtr[left.size+1].(*treeNonLeafNode).parent = left
			left.subPtr[left.size+1].(*treeNonLeafNode).parentIndex = left.size
		case *treeLeafNode:
			left.subPtr[left.size+1].(*treeLeafNode).parent = left
			left.subPtr[left.size+1].(*treeLeafNode).parentIndex = left.size
		}
		left.size++
	}
	nonLeaf.link.deleteSelf()
	tree.deleteFromNonLeaf(nonLeaf.parent, nonLeaf.parentIndex)
}

func (tree *bPlusTree) mergeRightNonLeaf(nonLeaf, right *treeNonLeafNode) {
	for i := 0; i < right.size; i++ {
		nonLeaf.keys[nonLeaf.size] = right.keys[i]
		nonLeaf.subPtr[nonLeaf.size+1] = right.subPtr[i+1]
		switch nonLeaf.subPtr[nonLeaf.size+1].(type) {
		case *treeNonLeafNode:
			nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parent = nonLeaf
			nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parentIndex = nonLeaf.size
		case *treeLeafNode:
			nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parent = nonLeaf
			nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parentIndex = nonLeaf.size
		}
		nonLeaf.size++
	}
	right.link.deleteSelf()
	tree.deleteFromNonLeaf(right.parent, right.parentIndex)
}

func shiftFromRightNonLeaf(nonLeaf, right *treeNonLeafNode) {
	nonLeaf.keys[nonLeaf.size] = right.keys[0]
	nonLeaf.subPtr[nonLeaf.size+1] = right.subPtr[1]
	switch nonLeaf.subPtr[nonLeaf.size+1].(type) {
	case *treeNonLeafNode:
		nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parent = nonLeaf
		nonLeaf.subPtr[nonLeaf.size+1].(*treeNonLeafNode).parentIndex = nonLeaf.size
	case *treeLeafNode:
		nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parent = nonLeaf
		nonLeaf.subPtr[nonLeaf.size+1].(*treeLeafNode).parentIndex = nonLeaf.size
	}
	nonLeaf.size++
	right.size--
	for i := 0; i < right.size; i++ {
		right.keys[i] = right.keys[i+1]
		right.subPtr[i+1] = right.subPtr[i+2]
		switch right.subPtr[i+1].(type) {
		case *treeLeafNode:
			right.subPtr[i+1].(*treeLeafNode).parentIndex--
		case *treeNonLeafNode:
			right.subPtr[i+1].(*treeNonLeafNode).parentIndex--
		}
	}
	replaceRecursive(right.parent, right.parentIndex, right.keys[0])
}

func shiftFromLeftNonLeaf(left, nonLeaf *treeNonLeafNode) {
	for i := nonLeaf.size; i > 0; i-- {
		nonLeaf.keys[i] = nonLeaf.keys[i-1]
		nonLeaf.subPtr[i+1] = nonLeaf.subPtr[i]
		switch nonLeaf.subPtr[i].(type) {
		case *treeNonLeafNode:
			nonLeaf.subPtr[i].(*treeNonLeafNode).parentIndex++
		case *treeLeafNode:
			nonLeaf.subPtr[i+1].(*treeLeafNode).parentIndex++
		}
	}
	nonLeaf.size++
	left.size--
	nonLeaf.keys[0] = left.keys[left.size]
	nonLeaf.subPtr[1] = left.subPtr[left.size+1]
	switch nonLeaf.subPtr[1].(type) {
	case *treeLeafNode:
		nonLeaf.subPtr[1].(*treeLeafNode).parentIndex = 0
		nonLeaf.subPtr[1].(*treeLeafNode).parent = nonLeaf
	case *treeNonLeafNode:
		nonLeaf.subPtr[1].(*treeNonLeafNode).parentIndex = 0
		nonLeaf.subPtr[1].(*treeNonLeafNode).parent = nonLeaf
	}
	replaceRecursive(nonLeaf.parent, nonLeaf.parentIndex, nonLeaf.keys[0])
}

func simpleDeleteFromNonLeaf(nonLeaf *treeNonLeafNode, delete int) {
	nonLeaf.size--
	for i := delete; i < nonLeaf.size; i++ {
		nonLeaf.keys[i] = nonLeaf.keys[i+1]
		nonLeaf.subPtr[i+1] = nonLeaf.subPtr[i+2]
		switch nonLeaf.subPtr[i+1].(type) {
		case *treeLeafNode:
			nonLeaf.subPtr[i+1].(*treeLeafNode).parentIndex = i
		case *treeNonLeafNode:
			nonLeaf.subPtr[i+1].(*treeNonLeafNode).parentIndex = i
		}
	}
	if delete == 0 && nonLeaf.parent != nil && nonLeaf.parentIndex != -1 {
		replaceRecursive(nonLeaf.parent, nonLeaf.parentIndex, nonLeaf.keys[0])
	}
}
