package collection

/*
* @see     Set
* @see     List
* @see     Map
* @see     SortedSet
* @see     SortedMap
* @see     HashSet
* @see     TreeSet
* @see     ArrayList
* @see     LinkedList
* @see     Vector
* @see     Collections
* @see     Arrays
* @see     AbstractCollection
 */

type Collection interface {
	/*
			查看当前的集合的size大小
		    index 0
	*/
	Size() int

	/*
			判断当前集合是否为空
		    如果为空 return true
		    否则 return false
	*/
	IsEmpty() bool
	/*
		查询当前接口是否为空

	*/
	Contains(o interface{}) bool

	/*
		集合添加元素
	*/
	Add(o interface{}) bool

	/*
		删除元素
	*/
	Remove(index int) bool

	hashCode() int

	equals(o interface{}) bool

}
