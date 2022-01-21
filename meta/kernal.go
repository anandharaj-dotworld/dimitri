package meta

import (
	"dimitri/types"
	"dimitri/utils"
	"syscall"
)

var uname syscall.Utsname
var kernalstruct types.Kernal

func Kernal() types.Kernal {
	if err := syscall.Uname(&uname); err == nil {
		kernalstruct = types.Kernal{
			Sysname: utils.Int8ToStr(uname.Sysname[:]),
			Release: utils.Int8ToStr(uname.Release[:]),
			Version: utils.Int8ToStr(uname.Version[:]),
			Arch:    utils.Int8ToStr(uname.Machine[:]),
			Name:    utils.Int8ToStr(uname.Nodename[:]),
		}
	}
	return kernalstruct
}
