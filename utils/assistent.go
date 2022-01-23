package utils

import (
	"bufio"
	"bytes"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Int8ToStr(arr []int8) string {
	b := make([]byte, 0, len(arr))
	for _, v := range arr {
		if v == 0x00 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}

func IsInstalled(pkg string) bool {
	cmd, err := exec.Command("which", pkg).Output()
	var output bool
	if len(cmd) > 0 {
		output = true
	} else {
		output = false
	}
	if err != nil {
		log.Printf("error %s", err)
	}
	return output
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetInterface() string {
	dummy, dummy_b := exec.Command("iw", "dev"), new(bytes.Buffer)
	dummy.Stdout = dummy_b
	dummy.Run()
	dummy_s := bufio.NewScanner(dummy_b)
	dummy_t := dummy_s
	var interface_n string
	for dummy_s.Scan() {
		if strings.Contains(dummy_t.Text(), "Interface") {
			interface_n = strings.TrimSpace(dummy_s.Text())
		}
	}
	var interface_name string
	if len(interface_n) > 0 {
		interface_name = strings.Split(interface_n, " ")[1]
	} else {
		interface_name = "eth0"
	}

	return interface_name

}

func SystemMeta(path string) string {
	GetData, err := exec.Command("cat", path).Output()
	if err != nil {
		log.Println(err)
	}
	Data := strings.TrimSuffix(string(GetData), "\n")
	return Data
}

func GetProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return string(stdout)
}

func Checkdevbat() string {
	cc, bb := exec.Command("upower", "--show-info", "/org/freedesktop/UPower/devices/DisplayDevice"), new(bytes.Buffer)
	cc.Stdout = bb
	cc.Run()
	ss := bufio.NewScanner(bb)
	var power_supply string
	for ss.Scan() {
		if strings.Contains(ss.Text(), "power supply") {
			supply_state := strings.Split(ss.Text(), " ")
			if supply_state[12] == "yes" {
				power_supply = "found"
			} else {
				power_supply = "not found"
			}
		}
	}
	return power_supply
}

func VerifyAppInstalled(pkg string) bool {
	cmd, err := exec.Command("which", pkg).Output()
	var output bool
	if len(cmd) > 0 {
		output = true
	} else {
		output = false
	}
	if err != nil {
		log.Printf("error %s", err)
	}
	return output
}

func Installpkg(pkg string) string {
	cmd, err := exec.Command("apt-get", "install", pkg).Output()
	if err != nil {
		log.Printf("error %s", err)
	}
	output := string(cmd)
	return output
}

func BatTemp() (string, string) {
	cmd, _ := exec.Command("acpi", "-t").Output()
	adapter, err := exec.Command("acpi", "-a").Output()
	if err != nil {
		log.Printf("error %s", err)
	}
	output := string(cmd)
	adapter_output := string(adapter)
	spitadapter := strings.Split(adapter_output, " ")
	split := strings.Split(output, " ")
	return split[3], spitadapter[2]
}
