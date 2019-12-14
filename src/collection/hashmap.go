package collection

const (
	DEFAULT_INITAL_CAPACITY = 1 << 4

	DEFAULT_LOAD_FACTOR = 0.75
)

//链表结构里数据的数据类型 键值对
type kv struct {
	key   interface{}
	value interface{}
}

//链表结构
type linkNode struct {
	//节点数据
	hash uint64
	data kv
	//下一个节点
	next *linkNode
}

type HashMap struct {
	//HashMap木桶
	size    int
	buckets [DEFAULT_INITAL_CAPACITY]linkNode
}

func NewHashMap() *HashMap {
	return &HashMap{
	}
}

