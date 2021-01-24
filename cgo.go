package kakasi

// #cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/deps/include
// #cgo darwin LDFLAGS: -L${SRCDIR} -L${SRCDIR}/deps/darwin -lkakasi -lkakasi_wrapper
// #cgo linux LDFLAGS: -L${SRCDIR} -L${SRCDIR}/deps/linux -lkakasi
import "C"

import (
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/ysugimoto/go-kakasi/deps/darwin"
	_ "github.com/ysugimoto/go-kakasi/deps/include"
	_ "github.com/ysugimoto/go-kakasi/deps/linux"
)

// initalize: Resolve path to kanwadict
// Get runtime info of this file, and set environment variable that bundled path
func init() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	os.Setenv(
		"KANWADICTPATH",
		filepath.Join(dir, "./deps/share/kanwadict"),
	)
}
