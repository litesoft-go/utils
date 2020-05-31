package responsewriter

import (
	"net/http"

	// standard libs only above!

	"github.com/litesoft-go/utils/http/types"
)

//noinspection GoUnusedExportedFunction
func ReturnJSON(w http.ResponseWriter, statusCode int, msg string) {
	ReturnWithContentType(w, types.JSONcontentTypeWithCharSet, statusCode, msg)
}

//noinspection GoUnusedExportedFunction
func ReturnText(w http.ResponseWriter, statusCode int, msg string) {
	ReturnWithContentType(w, types.TEXTcontentTypeWithCharSet, statusCode, msg)
}

func ReturnWithContentType(w http.ResponseWriter, contentType string, statusCode int, msg string) {
	w.Header().Add("Content-Type", contentType)
	Return(w, statusCode, msg)
}

func Return(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(msg))
}
