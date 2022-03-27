package ring

import "fmt"

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

	ring.updateAssignments()

	return &ring
}

func (r *Ring) RegisterNode(id int) {
	node := Node{
		ID: id,
	}

	r.Nodes = append(r.Nodes, &node)
	r.updateAssignments()
}

func (r *Ring) ModifyFactor(by int) {
	for _, load := range r.Load {
		load.Factor = load.Factor + by
	}

	r.updateAssignments()
}

func (r *Ring) SetFactor(value int) {
	for _, load := range r.Load {
		load.Factor = value
	}

	r.updateAssignments()
}

func (r *Ring) updateAssignments() {
	res := make([]*Node, len(r.Nodes))
	for ix, no := range r.Nodes {
		res[ix] = &Node{
			ID: no.ID,
		}
	}

	allRanges := r.Load.getFactoredRanges()
	var i int
loop:
	for i < len(allRanges) {
		for _, node := range res {
			node.Load = append(node.Load, allRanges[i])

			if i == len(allRanges)-1 {
				break loop
			}

			i++
		}
	}

	r.Assignments = []*Node{}
	r.Assignments = append(r.Assignments, res...)
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
