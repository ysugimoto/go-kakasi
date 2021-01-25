package kakasi

// #include <stdlib.h>
// #include "kakasi_wrapper.h"
import "C"

import (
	"errors"
	"unsafe"
)

// kakasi transformation function
// This function calls kakasi C library via bridge function.
// Note that C bridge function cannot capture kakasi or runtime error
// So if transformation error is occurred, just return error as "transformation error".
// But, the bridge function outputs to stdout some error information (e.g. initialization error, etc...)
// Therefore you may see some logs in your stdout.
func Transform(text string, opts ...TransformOption) (string, error) {
	args := []*C.char{
		// Following three options are always required, due to Go <-> C I/O should be utf-8.
		C.CString("kakasi"),
		C.CString("-iutf-8"),
		C.CString("-outf-8"),
	}

	// Append user provided transform options
	for i := range opts {
		args = append(args, C.CString(opts[i]()))
	}

	words := C.CString(text)
	var retChar *C.char

	// Release pointer after function is end
	defer func() {
		C.free(unsafe.Pointer(words))
		for i := range args {
			C.free(unsafe.Pointer(args[i]))
		}
		if retChar != nil {
			C.free(unsafe.Pointer(retChar))
		}
	}()

	// Call C binding function
	retChar = C.Kakasi(C.int(len(args)), (**C.char)(&args[0]), words)
	if retChar == nil {
		return "", errors.New("Kakasi transformation error")
	}

	// Copy to Go string explicitly
	return C.GoString(retChar), nil
}
