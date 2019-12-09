package algraph

import "github.com/ClareChu/collection/src/graph/mgraph"

const MAX = 10

// 查看当前顶点是否被标记
type Mark struct {
	IsMark bool
}

func DFS(g *mgraph.MGraph, i int, marks []Mark) {
	marks[i].IsMark = true
	for j := 0; j < g.NumVertexes; j++ {
		if (g.EdgeTypes[i][j] == mgraph.EdgeType{Weight: MAX}) && !marks[j].IsMark {
			DFS(g, j, marks)
		}
	}
}

func DFSTraverse(g *mgraph.MGraph) {
	var marks []Mark
	for i := 0; i < g.NumVertexes; i++ {
		mark := Mark{IsMark: false}
		marks = append(marks, mark)
		DFS(g, i, marks)
	}
}
