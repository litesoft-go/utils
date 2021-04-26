package fs

import (
	"github.com/litesoft-go/utils/interfaces"
	"io"
)

func CloseQuietly(closer io.Closer) {
	if !interfaces.IsNil(closer) {
		_ = closer.Close()
	}
}

func EnsureClosed(closer io.Closer, err error) error {
	if err != nil {
		CloseQuietly(closer)
	} else if !interfaces.IsNil(closer) {
		err = closer.Close()
	}
	return err
}
