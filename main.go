package main

// #cgo CFLAGS: -I. -I./kakasi/kakasi/include
// #cgo LDFLAGS: -L. -L./kakasi/kakasi/lib -lkakasi -lkakasi_wrap
// #include <stdlib.h>
// #include "kakasi_wrap.h"
import "C"

import (
	"fmt"
	"strings"
	// "unsafe"
)

func main() {
	str := C.CString("群馬県")
	// defer C.free(unsafe.Pointer(str))

	args := []string{
		"-Ha",
		"-Ka",
		"-Ja",
		"-Ea",
		"-ka",
		"-i",
		"utf-8",
		"-o",
		"utf-8",
		"-u",
	}
	argc := C.int(len(args))
	argv := C.CString(strings.Join(args, " "))
	// defer C.free(unsafe.Pointer(argv))

	C.Kakasi(argc, argv, str)
	fmt.Println("finish:")
}
