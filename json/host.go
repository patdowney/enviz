package json

type Host struct {
	HostName    string
	DNSName     string
	IPAddress   string
	State       string
	ServiceRefs []string
}
