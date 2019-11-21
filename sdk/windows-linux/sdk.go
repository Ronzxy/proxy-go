package main

import (
	"C"

	sdk "github.com/skygangsta/proxy/sdk/android-ios"
)

//export Start
func Start(serviceID *C.char, serviceArgsStr *C.char) (errStr *C.char) {
	return C.CString(sdk.Start(C.GoString(serviceID), C.GoString(serviceArgsStr)))
}

//export Stop
func Stop(serviceID *C.char) {
	sdk.Stop(C.GoString(serviceID))
}

//export Version
func Version() (ver *C.char) {
	return C.CString(sdk.Version())
}

//export StartProfiling
func StartProfiling(storePath *C.char) {
	sdk.StartProfiling(C.GoString(storePath))
}

//export StopProfiling
func StopProfiling() {
	sdk.StopProfiling()
}

func main() {
}
