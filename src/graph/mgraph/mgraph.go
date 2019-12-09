package mgraph

import "errors"

/*

邻接矩阵*/

const MAX_INFINITY = 65535

type MGraph struct {
	VertexTypes []VertexType `json:"vertexTypes"` // 顶点表 数组
	EdgeTypes   [5][5]EdgeType `json:"edgeTypes"`   //图的边表 临接矩阵
	NumVertexes int          `json:"numVertexes"` // 定点数量
	NumEdge     int          `json:"numEdge"`     // 边数量
}

type VertexType struct {
	Id string `json:"id"`
}

type EdgeType struct {
	Weight int `json:"weight"`
}

func NewMGraph() (m *MGraph) {
	return &MGraph{
		VertexTypes: []VertexType{},
		EdgeTypes:   [5][5]EdgeType{},
		NumVertexes: 0,
		NumEdge:     0,
	}
}

// 添加节点
func (m *MGraph) AddVertex(vt VertexType) {
	m.VertexTypes = append(m.VertexTypes, vt)
	m.NumVertexes++
}

func (m *MGraph) AddEdgeType(et EdgeType, v, w int) (err error) {
	if v == w {
		return errors.New("start and end edge Not allowed the same as")
	} else if len(m.VertexTypes) < v || len(m.VertexTypes) < w {
		return errors.New("start and end edge Not allowed less than VertexTypes length")
	}
	m.EdgeTypes[v][w] = et
	return
}
