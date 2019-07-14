package BplusTree

/*
 * 用于链接B➕树的同级节点
 */
type link struct {
	pre  *link
	next *link
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
 * 把 add 添加到 l 前边
 */
func (l *link) addPre(add *link) {

}

/*
 * 把 add 添加到 l 后边
 */
func (l *link) addNext(add *link) {
	if l.next == nil {
		l.next = add
		add.pre = l
	} else {
		l.next.pre = add
		add.next = l.next
		add.pre = l
		l.next = add
	}
}
