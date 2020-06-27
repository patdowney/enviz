package resources

type Network struct {
	Name     string
	Hosts    []Host
	Segments []NetworkSegment
}

type NetworkSegment struct {
	Name  string
	Hosts []Host
}
