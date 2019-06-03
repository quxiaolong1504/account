package seqd

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/rand"
)

type seqd struct {
	nodes []*snowflake.Node
}

var Seqd *seqd

func init()  {
	Seqd = NewSeqd(100)
}

func Generate () snowflake.ID {
	return Seqd.nodes[rand.Intn(len(Seqd.nodes))].Generate()
}

func NewSeqd(nodeLimit int) *seqd {
	nodes := make([]*snowflake.Node, 0)
	for i := 0; i < nodeLimit; i++ {
		node, err := snowflake.NewNode(int64(i))
		if err != nil {
			fmt.Print(err)
		}
		nodes = append(nodes, node)
	}
	return &seqd{
		nodes: nodes,
	}
}

