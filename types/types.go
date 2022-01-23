package types

import "encoding/json"

type Kernal struct {
	Sysname string `json:"system_name"`
	Release string `json:"release"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
	Name    string `json:"name"`
}

type PairedBle struct {
	Name        string `json:"name"`
	Mac_address string `json:"mac"`
}

type Storage struct {
	TotalDisk            string `json:"total_internal_storage"`
	FreeDiskSpace        string `json:"available_internal_storage"`
	ExternalStorageFree  int    `json:"available_external_storage"`
	ExternalStorageTotal int    `json:"total_external_storage"`
}

type Memory struct {
	RamTotal     uint64 `json:"total_mem,omitempty"`
	RamAvailable uint64 `json:"avail_mem,omitempty"`
	RunTimeTotal uint64 `json:"run_time_total_memory,omitempty"`
	RunTimeMax   uint64 `json:"run_time_max_memory,omitempty"`
	Threshold    uint64 `json:"threshold,omitempty"`
	RunTimeFree  uint64 `json:"run_time_free_memory,omitempty"`
	LowMemory    bool   `json:"low_memory,omitempty"`
}

type Wifi struct {
	AvailableWifi   json.RawMessage `json:"available_wifi_lists"`
	WifiName        string          `json:"connected_ssid"`
	WifiMAC         string          `json:"mac_address"`
	NetworkID       string          `json:"networkId"`
	ConnectedWifiIP string          `json:"ip_address"`
	BSSID           string          `json:"bssid"`
	LinkSpeed       int             `json:"link_speed"`
}

type Battery struct {
	BatteryHealth      string `json:"health,omitempty"`
	BatteryLevel       int    `json:"level,omitempty"`
	BatteryPlugged     string `json:"plugged,omitempty"`
	BatteryStatus      string `json:"status,omitempty"`
	BatteryTechnology  string `json:"technology,omitempty"`
	BatteryTemperature string `json:"temperature,omitempty"`
	BatteryVoltage     string `json:"voltage,omitempty"`
}

type All struct {
	Kernal    Kernal    `json:"kernal"`
	PairedBle PairedBle `json:"paired_bt_devices"`
	Storage   Storage   `json:"storage"`
	Memory    Memory    `json:"memory"`
	Wifi      Wifi      `json:"wifi"`
	Battery   Battery   `json:"battery"`
}
