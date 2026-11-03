package search

import "github.com/masterstruct/Eunoia/internal/board"

type Node struct {
	// add fields, update MakeMove and UnmakeMove
	staticEval int32
	singular   board.Move
}

// headroom if accessing stack[ply+N]
type NodeStack [MaxPly + 8]Node

func newNodeStack() NodeStack {
	var stack NodeStack
	for i := range stack {
		stack[i] = Node{
			// set defaults
			staticEval: -INF,
			singular:   board.NullMove,
		}
	}
	return stack
}
