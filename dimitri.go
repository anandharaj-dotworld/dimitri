package main

import (
	"dimitri/go-figure"
	"dimitri/meta"
	"dimitri/utils"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	Version = "v1.0"
)

func main() {
	banner := figure.NewColorFigure("Dimitri", "isometric1", "green", true)
	banner.Print()
	if strings.TrimSuffix(utils.GetProcessOwner(), "\n") != "root" {
		fmt.Println("You are trying to run app as non root user, this application only run on root user!")
		fmt.Println("Make sure, the application running as root user!!!")
		return
	}
	if !contains(os.Args[1]) {
		fmt.Println("Option not found, use -h for more information")
		return
	}

	useVersion := flag.Bool("v", false, "Print Versions 😌")
	help := flag.Bool("h", false, "Print commands and options 😌")
	all := flag.Bool("a", false, "Print all meta data 😎")
	kernal := flag.Bool("k", false, "Print system kernal informations 😊")
	bluetooth := flag.Bool("bt", false, "Print Paired bluetooth devices 😇")
	storage := flag.Bool("s", false, "rint system storage informations 😄")
	memory := flag.Bool("m", false, "Print system memory informations 😅")
	wifi := flag.Bool("w", false, "Print wifi informations 😎")
	battery := flag.Bool("b", false, "print device battery details 😊")
	flag.Parse()
	if flag.NFlag() > 1 && flag.NFlag() < 1 {
		fmt.Println("This program needs exactly one argument")
		os.Exit(0)
	}

	if *useVersion {
		fmt.Println("Version: " + Version)
		return
	}

	if *help {
		fmt.Println("Dimitri " + Version + "\n" +
			"Usage: dimitri [OPTIONS]\n" +
			"OPTIONS:\n" +
			"-v: Print Versions 😌\n" +
			"-a: Print all meta data 😎\n" +
			"-k: Print system kernal informations 😊\n" +
			"-bt: Print Paired bluetooth devices 😇\n" +
			"-s: Print system storage informations 😄\n" +
			"-m: Print system memory informations 😅\n" +
			"-w: Print wifi informations 😎\n" +
			"-b: print device battery details 😊")
	}
	if *all {
		allString, _ := json.Marshal(meta.All())
		fmt.Println(string(allString))
	}

	if *kernal {
		kernalString, _ := json.Marshal(meta.Kernal())
		fmt.Println(string(kernalString))
	}
	if *bluetooth {
		bluetoothString, _ := json.Marshal(meta.GetBle())
		fmt.Println(string(bluetoothString))
	}
	if *storage {
		storageString, _ := json.Marshal(meta.GetStorage())
		fmt.Println(string(storageString))
	}
	if *memory {
		memoryString, _ := json.Marshal(meta.GetMemory())
		fmt.Println(string(memoryString))
	}
	if *wifi {
		wifiString, _ := json.Marshal(meta.GetWifi())
		fmt.Println(string(wifiString))
	}
	if *battery {
		batteryString, _ := json.Marshal(meta.GetBattery())
		fmt.Println(string(batteryString))
	}
}

func contains(str string) bool {
	s := []string{"-v", "-h", "-a", "-k", "-bt", "-s", "-m", "-w", "-b"}
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
