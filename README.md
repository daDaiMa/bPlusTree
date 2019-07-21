# B+树

#####   

**此B+树🌲，可以自定义key和存放的value**  

key可以重复,我打算把key重复的元素用另外一种数据结构连接在一起然后挂到B+树上，这样就保证了B+ 树节点中不会出现重复的key  


```go
//初始化一颗30阶的B+树 存放 Goods结构体，把price当作key, 也可以传入自定义的比价函数compareFunc 用于key的比较
type Goods struct {
		price int
		name  string
}  
tree := BplusTree.InitBPlusTree(30, compareFunc, Goods{}.price)
// 可以传入compareFunc 用于key的比较，上述例子使用 price当作key 
// 如果key是int float string 等基本类型 compareFunc可以为空（会采用基本类型的比较方式）
// 否则必须设置 compareFunc 

```

这个代码执行效率有点低，树的插入和删除过程中，数组有多次移动，因此可以优化。

然而这个优化也并没有什么卵用😳😳 这个连玩具都算不上

一开始就是为了写一个比较通用的，可以放在我下学期的数据结构课设上使用。