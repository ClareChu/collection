package algraph

type VertexType struct {
}

type EdgeType struct {
}

type EdgeNode struct {
	Adjvex int       `json:"adjvex"` // 临接点域 储存该节点对应的下标
	Weight EdgeType  `json:"weight"` //存储权重
	Next   *EdgeNode `json:"next"`   // 下一个接点的指针
}

type VertexNode struct {
	Data      VertexType `json:"data"`
	FirstNode *EdgeNode  `json:"firstNode"`
}

type AlGraph struct {
	AdjList     []VertexNode `json:"adjList"`
	NumVertexes int          `json:"numVertexes"` // 定点数量
	NumEdge     int          `json:"numEdge"`     // 边数量
}


