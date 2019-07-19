package BplusTree

func (tree *bPlusTree) Delete(key interface{}) {
	node := tree.root
	for {
		if leaf, ok := node.(*treeLeafNode); ok {
			tree.delete(key, leaf)
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

func (tree *bPlusTree) delete(key interface{}, leaf *treeLeafNode) {
	index := tree.binarySearch(leaf.keys, key, leaf.size)
	if index < 0 {
		return
	}
	simpleDelete(leaf, index)
	if leaf.parent == nil {
		return
	}

}

func simpleDelete(leaf *treeLeafNode, index int) {
	for i := index + 1; i < leaf.size; i++ {
		leaf.keys[i-1] = leaf.keys[i]
		leaf.data[i-1] = leaf.data[i-1]
	}
	leaf.size--
}
