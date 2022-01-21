package meta

import (
	"dimitri/types"
	"strconv"

	"github.com/shirou/gopsutil/disk"
)

var Storage types.Storage

func GetStorage() types.Storage {
	diskStat, _ := disk.Usage("/")
	TotalDisk := strconv.FormatUint(diskStat.Total, 10)
	FreeDiskSpace := strconv.FormatUint(diskStat.Free, 10)
	Storage = types.Storage{
		TotalDisk:            TotalDisk,
		FreeDiskSpace:        FreeDiskSpace,
		ExternalStorageTotal: 0,
		ExternalStorageFree:  0,
	}
	return Storage
}
