package ring

import "golang.org/x/exp/slices"

type Ring struct {
	Nodes []*Node
	Load  Partitions
}

func NewRing(nodeIDs ...int) *Ring {
	if len(nodeIDs) == 0 {
		return nil
	}

	var res []*Node
	nodes := createNodes(nodeIDs...)

	for _, node := range nodes {
		res = append(res, &node)
	}

	return &Ring{
		Nodes: res,
		Load:  NewPartitions(1),
	}
}

func createNodes(nodeIDs ...int) Nodes {
	if len(nodeIDs) == 0 {
		return nil
	}

	var res Nodes

	for _, id := range nodeIDs {
		res = append(res, Node{ID: id})
	}

	return res
}

func (r *Ring) ModifyFactor(by int) {
	for _, load := range r.Load {
		load.Factor = load.Factor + by
	}
}

func (r Ring) getAssignments() Nodes {
	var allRanges []Range

	for _, load := range r.Load {
		for i := 0; i < load.Factor; i++ {
			allRanges = append(allRanges, load.Range)
		}
	}

	res := make([]Node, len(r.Nodes))

loop:
	for i := 0; i < len(allRanges); i++ {
		for _, node := range res {
			if !slices.Contains(node.Load, allRanges[i]) {
				node.Load = append(node.Load, allRanges[i])
			}

			if i == len(allRanges)-1 {
				break loop
			}

			i++
		}
	}

	return res
}
