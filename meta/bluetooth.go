package meta

import (
	"dimitri/types"
	"dimitri/utils"
	"os/exec"
	"strings"
)

var Ble types.PairedBle

func GetBle() types.PairedBle {
	if utils.IsInstalled("bt-device") {
		ble_cmd, _ := exec.Command("bt-device", "-l").Output()
		split_out := strings.Split(string(ble_cmd), "\n")
		for _, data := range split_out {
			if !strings.Contains(data, "Added devices:") && len(data) > 0 {
				split := strings.Split(data, " (")
				mac := strings.TrimSuffix(split[1], ")")
				Ble = types.PairedBle{
					Mac_address: mac,
					Name:        split[0],
				}
			}
		}
	}
	return Ble
}
