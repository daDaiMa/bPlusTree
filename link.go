package BplusTree

/*
 * 用于链接B➕树的同级节点
 */
type link struct {
	pre  *link
	next *link
}

/*
 * 初始化link,自己的前后均指向自己
 */
func newLink() *link {
	l := &link{}
	l.pre = l
	l.next = l
	return l
}

/*
 * 把 add 添加到 l 前边
 */
func (l *link) addPre(add *link) {
	add.pre = l.pre
	l.pre.next = add
	add.next = l
	l.pre = add
}

/*
 * 把 add 添加到 l 后边
 */
func (l *link) addNext(add *link) {
	l.next.pre = add
	add.next = l.next
	l.next = add
	add.pre = l
}
