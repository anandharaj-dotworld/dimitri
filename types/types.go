package types

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

type All struct {
	Kernal    Kernal    `json:"kernal"`
	PairedBle PairedBle `json:"paired_bt_devices"`
	Storage   Storage   `json:"storage"`
	Memory    Memory    `json:"memory"`
}
