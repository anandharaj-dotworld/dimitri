package utils

import (
	"log"
	"os/exec"
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
