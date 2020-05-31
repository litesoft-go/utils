package request

import (
	"net/http"

	// standard libs only above!

	"github.com/litesoft-go/utils/ioutils"
)

//noinspection GoUnusedExportedFunction
func Drain(r *http.Request) {
	if r != nil {
		ioutils.DrainReadCloser(r.Body)
	}
}
