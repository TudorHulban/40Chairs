package ring

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/exp/slices"
)

type Node struct {
	Load []Range
	Sock string
	ID   int
}

type Nodes []*Node

func (n Nodes) WriteTo(w io.Writer) (int, error) {
	if len(n) == 0 {
		return 0, errors.New("no nodes to consider for writing to")
	}

	var res []string

	for _, node := range n {
		res = append(res, fmt.Sprintf("Node ID: %d, load: %v.", node.ID, node.Load))
	}

	res = append(res, "\n")

	return w.Write([]byte(strings.Join(res, "\n")))
}

func (n Nodes) getRanges(ids ...int) []Range {
	var res []Range

	if len(ids) == 0 {
		for _, node := range n {
			res = append(res, node.Load...)
		}
	} else {
		for _, node := range n {
			for _, id := range ids {
				if id == node.ID {
					res = append(res, node.Load...)
				}
			}
		}
	}

	slices.Sort[Range](res)

	return res
}

func (n Nodes) resetRanges() {
	for _, node := range n {
		*node = Node{ID: node.ID}
	}
}

func (n Nodes) getIDs() []int {
	var res []int

	for _, node := range n {
		res = append(res, node.ID)
	}

	return res
}
