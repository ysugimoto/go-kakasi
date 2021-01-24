package kakasi

// #cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/deps/include
// #cgo darwin LDFLAGS: -L${SRCDIR} -L${SRCDIR}/deps/darwin -lkakasi -lkakasi_wrapper
// #cgo linux LDFLAGS: -L${SRCDIR} -L${SRCDIR}/deps/linux -lkakasi
import "C"

import (
	_ "github.com/ysugimoto/go-kakasi/deps/darwin"
	_ "github.com/ysugimoto/go-kakasi/deps/include"
	_ "github.com/ysugimoto/go-kakasi/deps/linux"
)
