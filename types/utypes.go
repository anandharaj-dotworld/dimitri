package types

type Cell struct {
	MAC           string  `json:"bssid"`
	Encryption    string  `json:"capabilities"`
	Channel       int     `json:"channel_width"`
	Frequency     float32 `json:"frequency"`
	EncryptionKey bool    `json:"is_passpoint_network"`
	SignalLevel   int     `json:"level"`
	ESSID         string  `json:"ssid"`
	NetworkTime   int64   `json:"timestamp"`
	VenueName     string  `json:"venue_name"`
	Quality       int     `json:"signal_quality"`
}
