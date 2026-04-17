package xifty

/*
#cgo CFLAGS: -I${SRCDIR}/internal/include
#cgo LDFLAGS: -L${SRCDIR}/.xifty-core/target/debug -lxifty_ffi
#include "xifty.h"
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type ViewMode int32

const (
	ViewFull ViewMode = iota
	ViewRaw
	ViewInterpreted
	ViewNormalized
	ViewReport
)

func Version() string {
	return C.GoString(C.xifty_version())
}

func Probe(path string) (map[string]any, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	result := C.xifty_probe_json(cPath)
	return decodeResult(result)
}

func Extract(path string, view ViewMode) (map[string]any, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	result := C.xifty_extract_json(cPath, uint32(view))
	return decodeResult(result)
}

func decodeResult(result C.XiftyResult) (map[string]any, error) {
	defer C.xifty_free_buffer(result.output)
	defer C.xifty_free_buffer(result.error_message)

	if result.status != C.XIFTY_STATUS_CODE_SUCCESS {
		return nil, fmt.Errorf("xifty ffi error %d: %s", int(result.status), bufferString(result.error_message))
	}

	var parsed map[string]any
	if err := json.Unmarshal([]byte(bufferString(result.output)), &parsed); err != nil {
		return nil, err
	}
	return parsed, nil
}

func ExtractNamed(path string, view string) (map[string]any, error) {
	switch view {
	case "full":
		return Extract(path, ViewFull)
	case "raw":
		return Extract(path, ViewRaw)
	case "interpreted":
		return Extract(path, ViewInterpreted)
	case "normalized":
		return Extract(path, ViewNormalized)
	case "report":
		return Extract(path, ViewReport)
	default:
		return nil, fmt.Errorf("unsupported view: %s", view)
	}
}

func bufferString(buffer C.XiftyBuffer) string {
	if buffer.ptr == nil || buffer.len == 0 {
		return ""
	}
	return C.GoStringN((*C.char)(unsafe.Pointer(buffer.ptr)), C.int(buffer.len))
}
