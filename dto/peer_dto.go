package dto

// PeerDto PeerDto
type PeerDto struct {
	IpAddress                  string
	ProtocolVersion            int
	ConnectionTime             int64
	ConnectionStatus           string
	Inbound                    bool
	BufferedTransactionsCount  int
	BufferedBlocksCount        int
	BufferedAnnouncementsCount int
	RequestMetrics             []RequestMetric
	NodeVersion                string
}

// RequestMetric RequestMetric
type RequestMetric struct {
	RoundTripTime int64
	MethodName    string
	Info          string
	RequestTime   string
}
