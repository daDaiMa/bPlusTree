# B+树

就是想写一颗B+树，觉得很有意思    

此B+树，可以自定义key和存放的value  
key可以重复,我打算把key重复的元素用另外一种数据结构连接在一起然后挂到B+树上，这样就保证了B+ 树节点中不会出现重复的key  

目前就写了插入和查找，还没写删除  


//初始化一颗30阶的B+树 存放 Goods结构体，把price当作key, 也可以传入自定义的比价函数compareFunc 用于key的比较
type Goods struct {
		price int
		name  string
}  
tree := BplusTree.InitBPlusTree(30, compareFunc, Goods{}.price)

