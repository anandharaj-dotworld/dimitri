package main

import (
	"dimitri/go-figure"
	"dimitri/meta"
	"encoding/json"
	"flag"
	"fmt"
)

var (
	version = "v1.0"
)

func main() {
	banner := figure.NewColorFigure("Dimitri", "isometric1", "green", true)
	banner.Print()
	useVersion := flag.Bool("v", false, "display output")
	help := flag.Bool("h", false, "display output")
	all := flag.Bool("a", false, "display output")
	kernal := flag.Bool("k", false, "display output")
	bluetooth := flag.Bool("b", false, "display output")
	storage := flag.Bool("s", false, "display output")
	memory := flag.Bool("m", false, "display output")
	flag.Parse()
	if *useVersion {
		fmt.Println("Version: " + version)
		return
	}

	if *help {
		fmt.Println("Dimitri " + version + "\n" +
			"Usage: dimitri [OPTIONS]\n" +
			"OPTIONS:\n" +
			"-v: Print Versions\n" +
			"-a: Print all meta data" +
			"-k: Print system kernal informations\n" +
			"-b: Print Paired bluetooth devices\n" +
			"-s: Print system storage informations\n" +
			"-m: Print system memory informations")
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
}
