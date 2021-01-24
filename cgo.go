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

// Resolve path to kanwadict and itaijidict and set envronment path
// This envioment variables must be set before calling bridge function
// because kakasi is built with our local path and then these dict path is set as local path.
// Fortunately, kakasi can lookup dict files from environment and prior to lookup built path
// so we can resolve these dict files by setting environment variable like KANWADICTPATH and ITAIJIDICTPATH.
// In go, to resolve bundle path by using runtime package to get actual THIS filepath.
func init() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	os.Setenv(
		"KANWADICTPATH",
		filepath.Join(dir, "./deps/share/kanwadict"),
	)
	os.Setenv(
		"ITAIJIDICTPATH",
		filepath.Join(dir, "./deps/share/itaijidict"),
	)
}
