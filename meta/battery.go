package meta

import (
	"bufio"
	"bytes"
	"dimitri/types"
	"dimitri/utils"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	battery_level       int
	health              string
	battery_technology  string
	battery_voltage     string
	battery_temperature string
	battery_status      string
	battery_plugged     string
)

func GetBattery() types.Battery {
	if utils.Checkdevbat() == "found" {
		verify := utils.VerifyAppInstalled("acpi")
		if !verify {
			utils.Installpkg("acpi")
		}
		verify_sensors := utils.VerifyAppInstalled("sensors")
		if !verify_sensors {
			utils.Installpkg("lm-sensors")
		}

		files, _ := os.ReadDir("/sys/class/power_supply/")
		file_count := 0
		bat_lev := 0
		bat_full := 0
		for _, file := range files {
			if strings.HasPrefix(file.Name(), "BAT") {
				if _, err := os.Stat("/sys/class/power_supply/" + file.Name() + "/energy_now"); err == nil {
					file_count += 1
					getbl := utils.SystemMeta("/sys/class/power_supply/" + file.Name() + "/energy_now")
					getblfull := utils.SystemMeta("/sys/class/power_supply/" + file.Name() + "/energy_full")
					trimSbl, _ := strconv.Atoi(strings.TrimSuffix(getbl, "\n"))
					trimSblfull, _ := strconv.Atoi(strings.TrimSuffix(getblfull, "\n"))
					bat_lev += trimSbl
					bat_full += trimSblfull
				}
			}
		}

		if bat_lev == 0 {
			command, battery_bytes := exec.Command("upower", "--show-info", "/org/freedesktop/UPower/devices/DisplayDevice"), new(bytes.Buffer)
			command.Stdout = battery_bytes
			command.Run()
			battery_scanner := bufio.NewScanner(battery_bytes)
			tt := battery_scanner
			var level string
			for battery_scanner.Scan() {
				if strings.Contains(tt.Text(), "percentage") {
					level = battery_scanner.Text()
				}
			}
			re := regexp.MustCompile("[0-9]+")
			lvl := re.FindAllString(level, -1)

			battery_level, _ = strconv.Atoi(lvl[0])
		} else {
			batfullm := float64(bat_lev) / float64(bat_full)
			battery_level, _ = strconv.Atoi(fmt.Sprintf("%.0f", batfullm*100))
		}

		if battery_level > 90 {
			health = "GOOD"
		} else {
			health = "NORMAL"
		}
		battery_technology = utils.SystemMeta("/sys/class/power_supply/BAT0/technology")
		battery_voltage_t, _ := strconv.Atoi(utils.SystemMeta("/sys/class/power_supply/BAT0/voltage_now"))
		battery_voltage = strconv.Itoa(battery_voltage_t)[:2]
		_, battery_statuss := utils.BatTemp()
		bat_status := strings.TrimSuffix(battery_statuss, "\n")

		if bat_status == "on-line" {
			battery_status = "CHARGING"
			battery_plugged = "AC"
		} else if bat_status == "off-line" {
			battery_status = "DISCHARGING"
			battery_plugged = "DC"
		}

		sensors, _ := utils.NewFromSystem()
		var battery_trim string
		for chip := range sensors.Chips {
			for key, value := range sensors.Chips[chip] {
				if key == "Sensor 1" {
					battery_trim = value
				}
			}
		}
		battery_temperature = strings.TrimPrefix(strings.TrimSuffix(strings.Split(battery_trim, " ")[0], "°C"), "+")
	}
	return types.Battery{
		BatteryHealth:      health,
		BatteryLevel:       battery_level,
		BatteryPlugged:     battery_plugged,
		BatteryStatus:      battery_status,
		BatteryTechnology:  battery_technology,
		BatteryTemperature: battery_temperature,
		BatteryVoltage:     battery_voltage,
	}
}
