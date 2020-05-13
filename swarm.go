package config

type SwarmConfig struct {
	AddrFilters             []string
	DisableBandwidthMetrics bool
	DisableNatPortMap       bool
	DisableRelay            bool
	EnableRelayHop          bool
	SwarmKey                string

	// autorelay functionality
	// if true, then the libp2p host will be constructed with autorelay functionality.
	EnableAutoRelay bool

	ConnMgr ConnMgr
}

// ConnMgr defines configuration options for the libp2p connection manager
type ConnMgr struct {
	Type        string
	LowWater    int
	HighWater   int
	GracePeriod string
}
