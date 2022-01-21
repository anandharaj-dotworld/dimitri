package meta

import "dimitri/types"

func All() types.All {
	allType := types.All{
		Kernal:    Kernal(),
		PairedBle: GetBle(),
		Storage:   GetStorage(),
		Memory:    GetMemory(),
	}
	return allType
}
