package fs

import "io"

func CloseQuietly(closer io.Closer) {
	_ = closer.Close()
}

func EnsureClosed(closer io.Closer, err error) error {
	if err != nil {
		CloseQuietly(closer)
	} else {
		err = closer.Close()
	}
	return err
}
