package ioutils

import (
	"io"
	"io/ioutil"

	// standard libs only above!

	"github.com/litesoft-go/utils/interfaces"
)

func Close(in io.Closer) {
	_ = in.Close() // for http Body(s) Per Dave Cheney 2017 - auto drains! https://forum.golangbridge.org/t/do-i-need-to-read-the-body-before-close-it/5594
}

//noinspection GoUnusedExportedFunction
func DrainReadCloser(in io.ReadCloser) {
	if !interfaces.IsNil(in) {
		defer Close(in)
		DrainReader(in)
	}
}

func DrainReader(in io.Reader) {
	if !interfaces.IsNil(in) {
		_, _ = io.Copy(ioutil.Discard, in)
	}
}

//noinspection GoUnusedExportedFunction
func ReadAllClose(in io.ReadCloser) (bytes []byte, err error) {
	if !interfaces.IsNil(in) {
		defer Close(in)
		bytes, err = ioutil.ReadAll(in)
	}
	return
}

//noinspection GoUnusedExportedFunction
func WriteAll(bytes []byte, out io.Writer) error {
	for from := 0; from < len(bytes); {
		wrote, err := out.Write(bytes[from:])
		if err != nil {
			return err
		}
		from += wrote
	}
	return nil
}
