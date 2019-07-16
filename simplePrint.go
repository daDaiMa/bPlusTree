package BplusTree

import "fmt"

func (tree *bPlusTree) PrintSimply() {
	queue := make([]interface{}, 0)
	queue = append(queue, tree.root)
	child := make([]interface{}, 0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if leaf, ok := node.(*treeLeafNode); ok {
			for i := 0; i < leaf.size; i++ {
				fmt.Print(leaf.keys[i], " ")
			}
			fmt.Printf("\t")
		} else {
			if node == nil {
				continue
			}
			nonLeaf := node.(*treeNonLeafNode)
			for i := 0; i < nonLeaf.size; i++ {
				fmt.Print(nonLeaf.keys[i], " ")
			}
			fmt.Printf("\t")
			for i := 0; i <= nonLeaf.size; i++ {
				child = append(child, nonLeaf.subPtr[i])
			}
		}

		if len(queue) == 0 {
			fmt.Printf("\n")
			if len(child) == 0 {
				return
			}

			queue = child
			child = make([]interface{}, 0)
		}
	}
}
