package ring

import "golang.org/x/exp/slices"

type Range string

type Partition struct {
	Range  Range
	Factor int // redundancy factor for partition
}

type Partitions []Partition

func NewPartitions(factor int) Partitions {
	return []Partition{
		Partition{"0", factor},
		Partition{"1", factor},
		Partition{"2", factor},
		Partition{"3", factor},
		Partition{"4", factor},
		Partition{"5", factor},
		Partition{"6", factor},
		Partition{"7", factor},
		Partition{"8", factor},
		Partition{"9", factor},
	}
}

func (p Partitions) getRanges() []string {
	var res []string

	for _, partition := range p {
		res = append(res, string(partition.Range))
	}

	slices.Sort[string](res)

	return res
}
