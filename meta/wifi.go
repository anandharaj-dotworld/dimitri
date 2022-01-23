package meta

import (
	"dimitri/types"
	"dimitri/utils"
	wlist "dimitri/wifi_lib"
	"encoding/json"
	"net"
	"os/exec"
	"strings"
)

var wifiTypes types.Wifi

func GetWifi() types.Wifi {
	wname, _ := exec.Command("iwgetid", "--raw").Output()
	wifi_name := strings.TrimSuffix(string(wname), "\n")
	ConnectedWifiIP := utils.GetLocalIP()
	var cell []types.Cell
	intface := utils.GetInterface()
	cell, _ = wlist.Scan(intface)
	available_wifi_device, _ := json.Marshal(cell)
	var link_speed int
	var bssid_id string
	for _, value := range cell {
		if wifi_name == value.ESSID {
			link_speed = value.Quality
			bssid_id = value.MAC
		}
	}

	ifas, _ := net.Interfaces()
	var wifi_mac string
	for _, ifa := range ifas {
		if ifa.Name == intface {
			wifi_mac = ifa.HardwareAddr.String()
		}
	}
	NetworkID := utils.SystemMeta("/sys/class/net/" + intface + "/netdev_group")
	wifiTypes = types.Wifi{
		AvailableWifi:   available_wifi_device,
		WifiName:        wifi_name,
		WifiMAC:         wifi_mac,
		NetworkID:       NetworkID,
		ConnectedWifiIP: ConnectedWifiIP,
		BSSID:           bssid_id,
		LinkSpeed:       link_speed,
	}
	return wifiTypes
}
