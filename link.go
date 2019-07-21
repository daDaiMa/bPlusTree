package BplusTree

/*
 * 用于链接B➕树的同级节点
 */
type link struct {
	pre  **link
	next **link
}

/*
 * 初始化link, 感觉不需要循环链表
 */
func newLink() *link {
	l := &link{}
	l.pre = nil
	l.next = nil
	return l
}

/*
 * 把 add 添加到 l 后边
 */
func addNext(ori, add **link) {
	if (*ori).next == nil {
		(*ori).next = add
		(*add).pre = ori
	} else {
		(*(*ori).next).pre = add
		(*add).next = (*ori).next
		(*add).pre = ori
		(*ori).next = add
	}
}

func (l *link) deleteSelf() {
	if l.pre != nil {
		(*l.pre).next = l.next
	}
	if l.next != nil {
		(*l.next).pre = l.pre
	}
}

