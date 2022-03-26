package ring

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Node struct {
	Load []Range
	ID   int
}

type Nodes []Node

func (n Nodes) WriteTo(w io.Writer) (int, error) {
	if len(n) == 0 {
		return 0, errors.New("no nodes to consider for writing to")
	}

	var res []string

	for _, node := range n {
		res = append(res, fmt.Sprintf("Node ID: %d, load: %v.", node.ID, node.Load))
	}

	return w.Write([]byte(strings.Join(res, "\n")))
}
