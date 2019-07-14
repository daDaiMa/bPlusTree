package BplusTree

func (tree *bPlusTree) Search(key interface{}) interface{} {
	node := tree.root
	for {
		if leaf, ok := node.(*treeLeafNode); ok {
			index := tree.binarySearch(leaf.keys, key, leaf.size)
			if index >= 0 {
				return leaf.data[index]
			} else {
				return nil
			}
		} else {
			index := tree.binarySearch(node.(*treeNonLeafNode).keys, key, node.(*treeNonLeafNode).size)
			if index >= 0 {
				node = node.(*treeNonLeafNode).subPtr[index+1]
			} else {
				node = node.(*treeNonLeafNode).subPtr[-index-1]
			}
		}
	}
}
