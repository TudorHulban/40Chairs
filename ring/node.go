package ring

type Node struct {
	Load   []Range
	Sock   string
	ID     int
	RootID int
}
