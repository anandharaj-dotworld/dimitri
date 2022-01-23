package wlist

import (
	"dimitri/types"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	newCellRegexp = regexp.MustCompile(`^Cell\s+(?P<cell_number>.+)\s+-\s+Address:\s(?P<mac>.+)$`)
	regxp         [7]*regexp.Regexp
)

func init() {
	// precompile regexp
	regxp = [7]*regexp.Regexp{
		regexp.MustCompile(`^ESSID:\"(?P<essid>.*)\"$`),
		regexp.MustCompile(`^Mode:(?P<mode>.+)$`),
		regexp.MustCompile(`^Frequency:(?P<frequency>[\d.]+) (?P<frequency_units>.+) \(Channel (?P<channel>\d+)\)$`),
		regexp.MustCompile(`^Encryption key:(?P<encryption_key>.+)$`),
		regexp.MustCompile(`^IE:\ WPA\ Version\ (?P<wpa>.+)$`),
		regexp.MustCompile(`^IE:\ IEEE\ 802\.11i/WPA2\ Version\ (?P<wpa2>)$`),
		regexp.MustCompile(`^Quality=(?P<signal_quality>\d+)/(?P<signal_total>\d+)\s+Signal level=(?P<signal_level>.+) d.+$`),
	}
}

func Scan(interfaceName string) ([]types.Cell, error) {
	// execute iwlist for scanning wireless networks
	cmd := exec.Command("iwlist", interfaceName, "scan")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// parse fetched result
	return parse(string(out))
}

func parse(input string) (cells []types.Cell, err error) {
	lines := strings.Split(input, "\n")

	var cell *types.Cell
	var wg sync.WaitGroup
	var m sync.Mutex
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// check new cell value
		if cellValues := newCellRegexp.FindStringSubmatch(line); len(cellValues) > 0 {
			cells = append(cells, types.Cell{
				MAC: cellValues[2],
			})
			cell = &cells[len(cells)-1]

			continue
		}

		// compare lines to regexps
		wg.Add(len(regxp))
		for _, reg := range regxp {
			go compare(line, &wg, &m, cell, reg)
		}
		wg.Wait()
	}

	return
}

func compare(line string, wg *sync.WaitGroup, m *sync.Mutex, cell *types.Cell, reg *regexp.Regexp) {
	defer wg.Done()
	if values := reg.FindStringSubmatch(line); len(values) > 0 {
		keys := reg.SubexpNames()

		m.Lock()

		for i := 1; i < len(keys); i++ {
			switch keys[i] {
			case "essid":
				cell.ESSID = values[i]
				now := time.Now()
				sec := now.Unix()
				cell.NetworkTime = sec
				cell.VenueName = ""
			case "frequency":
				if frequency, err := strconv.ParseFloat(values[i], 32); err == nil {
					cell.Frequency = float32(frequency)
				}
			case "channel":
				if channel, err := strconv.ParseInt(values[i], 10, 32); err == nil {
					cell.Channel = int(channel)
				}
			case "encryption_key":
				if cell.EncryptionKey = values[i] == "on"; cell.EncryptionKey {
					cell.Encryption = "wep"
				} else {
					cell.Encryption = "off"
				}
			case "wpa":
				cell.Encryption = "wpa"
			case "wpa2":
				cell.Encryption = "wpa2"
			case "signal_level":
				if level, err := strconv.ParseInt(values[i], 10, 32); err == nil {
					cell.SignalLevel = int(level)

				}
			case "signal_quality":
				if quality, err := strconv.ParseInt(values[i], 10, 32); err == nil {
					cell.Quality = int(quality)
				}
			}
		}

		m.Unlock()
	}
}
