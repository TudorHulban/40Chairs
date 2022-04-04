package ring

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Ring struct {
	Nodes       Nodes
	Assignments Nodes
	Load        Partitions
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

func NewRing(nodeIDs ...int) *Ring {
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
		Load:  NewPartitions(1),
	}

	ring.renewAssignments()

	return &ring
}

func (r *Ring) RegisterNode(nodeID int) {
	node := Node{
		ID: nodeID,
	}

	r.Nodes = append(r.Nodes, &node)
	r.renewAssignments()
}

func (r *Ring) UnRegisterNode(nodeID int) {
	if len(r.Nodes) == 0 {
		return
	}

	var nodeRanges []Range

	for ix, node := range r.Nodes {
		if nodeID == node.ID {
			nodeRanges = append(nodeRanges, r.Assignments[ix].Load...)

			copy(r.Nodes[ix:], r.Nodes[ix+1:])
			r.Nodes = r.Nodes[:len(r.Nodes)-1]

			break
		}
	}

	r.redistributeAssignmentsFrom(nodeID, nodeRanges)
}

func (r *Ring) ModifyFactor(by int) {
	for _, load := range r.Load {
		load.Factor = load.Factor + by
	}

	r.renewAssignments()
}

func (r *Ring) SetFactor(value int) {
	for _, load := range r.Load {
		load.Factor = value
	}

	r.renewAssignments()
}

func (r *Ring) resetAssignments() {
	r.Assignments = make([]*Node, len(r.Nodes))

	for ix, no := range r.Nodes {
		r.Assignments[ix] = &Node{
			ID: no.ID,
		}
	}
}

func (r *Ring) sortAssignments() {
	for _, assignment := range r.Assignments {
		slices.Sort[Range](assignment.Load)
	}
}

func (r *Ring) renewAssignments() {
	r.resetAssignments()

	allRanges := r.Load.getFactoredRanges()
	var i int

loop:
	for i < len(allRanges) {
		for _, node := range r.Assignments {
			node.Load = append(node.Load, allRanges[i])

			if i == len(allRanges)-1 {
				break loop
			}

			i++
		}
	}
}

func (r *Ring) redistributeAssignmentsFrom(nodeID int, ranges []Range) error {
	if slices.Contains[int](r.Nodes.getIDs(), nodeID) {
		return fmt.Errorf("redistribution cannot proceed as node with ID: %d was not removed from ring", nodeID)
	}

	var i int
	var ix int

	fmt.Println("ranges:", ranges)

loop:
	for i < len(ranges) {
		for j, node := range r.Assignments {
			if node.ID == nodeID {
				ix = j
				continue
			}

			if slices.Index[Range](node.Load, ranges[i]) == -1 {
				node.Load = append(node.Load, ranges[i])

				i++
			}

			if i == len(ranges) {
				break loop
			}
		}
	}

	copy(r.Assignments[ix:], r.Assignments[ix+1:])
	r.Assignments = r.Assignments[:len(r.Assignments)-1]

	r.sortAssignments()

	return nil
}

func (r *Ring) verifyAssignments() error {
	var assignedRanges []Range

	for _, node := range r.Assignments {
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
