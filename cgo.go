package kakasi

// #cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/deps/include
// #cgo darwin LDFLAGS: -L${SRCDIR} -L${SRCDIR}/deps/darwin -lkakasi -lkakasi_wrapper
// #cgo linux LDFLAGS: -L${SRCDIR} -L${SRCDIR}/deps/linux -lkakasi
// #cgo LDFLAGS: -Wl,-rpath,"${SRCDIR}"
import "C"

import (
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/rakyll/statik/fs"

	_ "github.com/ysugimoto/go-kakasi/deps/darwin"
	_ "github.com/ysugimoto/go-kakasi/deps/include"
	_ "github.com/ysugimoto/go-kakasi/deps/linux"
	_ "github.com/ysugimoto/go-kakasi/dict"
)

const (
	kanwadictFile  = "/tmp/kanwadict"
	itaijidictFile = "/tmp/itaijidict"
)

// Resolve path to kanwadict and itaijidict and set envronment path.
// Envioment variables must be set before calling bridge function because kakasi is built with our local path
// and then these dict path is set as local path.
// Fortunately, kakasi can lookup dict files from environment and prior to lookup built path
// so we can resolve these dict files by setting environment variable like KANWADICTPATH and ITAIJIDICTPATH.
// But dict files are loaded dynamically in C runtime and then these cannot resolve path on compiled binary,
// Therefor, dict files should extract bundled data and output to `/tmp/[dict file]`
// and set environment variables as that file path in initialize phase.
func init() {
	dict, err := fs.New()
	if err != nil {
		panic(err)
	}

	if err := putDictFile(dict, kanwadictFile, "/kanwadict"); err != nil {
		panic(err)
	}
	os.Setenv("KANWADICTPATH", kanwadictFile)

	if err := putDictFile(dict, itaijidictFile, "/itaijidict"); err != nil {
		panic(err)
	}
	os.Setenv("ITAIJIDICTPATH", itaijidictFile)
}

func putDictFile(d http.FileSystem, dest, src string) error {
	// Regarding dictfile already exists
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	srcFp, err := d.Open(src)
	if err != nil {
		return fmt.Errorf("Failed to open source file: %s, %w", src, err)
	}
	defer srcFp.Close()

	destFp, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open destination file: %s, %w", src, err)
	}
	defer destFp.Close()

	if _, err := io.Copy(destFp, srcFp); err != nil {
		return fmt.Errorf("Failed to copy file: %w", err)
	}
	return nil
}
