package meta

import (
	"bufio"
	"dimitri/types"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var memTypes types.Memory

func GetMemory() types.Memory {
	meminfo := &MemInfo{}
	meminfo.Update()
	RamTotal := meminfo.Total()
	RamFree := meminfo.Available()
	memTypes = types.Memory{
		RamTotal:     RamTotal,
		RamAvailable: RamFree,
	}
	return memTypes
}

type MemInfo map[string]uint64

func (m *MemInfo) Update() error {
	var err error

	path := filepath.Join("/proc/meminfo")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		n := strings.Index(text, ":")
		if n == -1 {
			continue
		}

		key := text[:n] // metric
		data := strings.Split(strings.Trim(text[(n+1):], " "), " ")
		if len(data) == 1 {
			value, err := strconv.ParseUint(data[0], 10, 64)
			if err != nil {
				continue
			}
			(*m)[key] = value
		} else if len(data) == 2 {
			if data[1] == "kB" {
				value, err := strconv.ParseUint(data[0], 10, 64)
				if err != nil {
					continue
				}
				(*m)[key] = value * 1024
			}
		}

	}
	return nil

}

func (m *MemInfo) Total() uint64 {
	return (*m)["MemTotal"]
}

func (m *MemInfo) Available() uint64 {
	d := *m
	return d["MemFree"] + d["Buffers"] + d["Cached"]
}

func (m *MemInfo) Used() uint64 {
	return m.Total() - m.Available()
}
func (m *MemInfo) Active() uint64 {
	d := *m
	return d["Active"]
}
func (m *MemInfo) InActive() uint64 {
	d := *m
	return d["Inactive"]
}

func (m *MemInfo) Swap() int {
	total := (*m)["SwapTotal"]
	free := (*m)["SwapFree"]
	if total == 0 {
		return 0
	}
	return int((100 * (total - free)) / total)
}
