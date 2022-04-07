package comms

type IComms interface {
	// AnnounceToTheJoined is to be used immediately after joining.
	// It would announce the joining node to the one pointed up in the start command.
	// The return should be the topology including the root node.
	AnnounceToTheJoined(sock string) string
	PingNode(sock string) error

	// SendRangesForNode would be invoked:
	// a. when node is joining receiving load ranges from root
	// b. when there is a redistribution of load
	SendRangesForNode(sock string, ranges []string)
}
