package ring

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Ring struct {
	Nodes Nodes
	Load  Partitions
}

func newNodes(nodeIDs ...int) Nodes {
	if len(nodeIDs) == 0 {
		return nil
	}

	var res Nodes

	for _, id := range nodeIDs {
		res = append(res, &Node{ID: id})
	}

	return res
}

func NewRing(factor int, nodeIDs ...int) *Ring {
	if len(nodeIDs) == 0 {
		return nil
	}

	var res []*Node
	nodes := newNodes(nodeIDs...)

	for _, node := range nodes {
		res = append(res, &Node{ID: node.ID})
	}

	ring := Ring{
		Nodes: res,
		Load:  NewPartitions(factor),
	}

	ring.resetAssignments()

	return &ring
}

func (r *Ring) RegisterNode(nodeID int) {
	node := Node{
		ID: nodeID,
	}

	r.Nodes = append(r.Nodes, &node)
	r.resetAssignments()
}

func (r *Ring) UnRegisterNode(nodeID int) {
	if len(r.Nodes) == 0 {
		return
	}

	var nodeRanges []Range

	for ix, node := range r.Nodes {
		if nodeID == node.ID {
			nodeRanges = append(nodeRanges, r.Nodes[ix].Load...)

			copy(r.Nodes[ix:], r.Nodes[ix+1:])
			r.Nodes = r.Nodes[:len(r.Nodes)-1]

			break
		}
	}

	r.redistributeLoad(nodeRanges)
}

func (r *Ring) ModifyFactor(by int) {
	for _, load := range r.Load {
		load.Factor = load.Factor + by
	}

	r.resetAssignments()
}

func (r *Ring) SetFactor(value int) {
	for _, load := range r.Load {
		load.Factor = value
	}

	r.resetAssignments()
}

func (r *Ring) sortRanges() {
	for _, node := range r.Nodes {
		slices.Sort[Range](node.Load)
	}
}

func (r *Ring) resetAssignments() {
	r.Nodes.resetRanges()

	allRanges := r.Load.getFactoredRanges()
	var i int

loop:
	for i < len(allRanges) {
		for _, node := range r.Nodes {
			node.Load = append(node.Load, allRanges[i])

			if i == len(allRanges)-1 {
				break loop
			}

			i++
		}
	}
}

func (r *Ring) redistributeLoad(ranges []Range) {
	var i int

loop:
	for i < len(ranges) {
		for _, node := range r.Nodes {
			if slices.Index[Range](node.Load, ranges[i]) == -1 {
				node.Load = append(node.Load, ranges[i])

				i++
			}

			if i == len(ranges) {
				break loop
			}
		}
	}

	r.sortRanges()
}

func (r *Ring) verifyAssignments() error {
	var assignedRanges []Range

	for _, node := range r.Nodes {
		assignedRanges = append(assignedRanges, node.Load...)
	}

	occ := occurences[Range](assignedRanges)

	for _, partition := range r.Load {
		if _, exists := occ[partition.Range]; !exists {
			return fmt.Errorf("not all partitions were mapped. missing '%v'", partition.Range)
		}

		if partition.Factor < occ[partition.Range] {
			return fmt.Errorf("for range '%s' factor is only %d versus required %d", partition.Range, occ[partition.Range], partition.Factor)
		}
	}

	return nil
}
