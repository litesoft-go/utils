package proxy

import (
	"io"
	"net/http"

	// standard libs only above!

	"github.com/litesoft-go/utils/http/method"
)

type BodiedClientFunc func(client *http.Client, url, contentType string, body io.Reader) (resp *http.Response, err error)

type NoBodyClientFunc func(client *http.Client, newURL string) (resp *http.Response, err error)

//noinspection GoUnusedExportedFunction
func POST(client *http.Client, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return bodied(client, url, contentType, body, method.POST)
}

//noinspection GoUnusedExportedFunction
func PUT(client *http.Client, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return bodied(client, url, contentType, body, method.PUT)
}

//noinspection GoUnusedExportedFunction
func PATCH(client *http.Client, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return bodied(client, url, contentType, body, method.PATCH)
}

//noinspection GoUnusedExportedFunction
func GET(client *http.Client, url string) (resp *http.Response, err error) {
	return noBody(client, url, method.GET)
}

//noinspection GoUnusedExportedFunction
func HEAD(client *http.Client, url string) (resp *http.Response, err error) {
	return noBody(client, url, method.HEAD)
}

//noinspection GoUnusedExportedFunction
func DELETE(client *http.Client, url string) (resp *http.Response, err error) {
	return noBody(client, url, method.DELETE)
}

//noinspection GoUnusedExportedFunction
func CONNECT(client *http.Client, url string) (resp *http.Response, err error) {
	return noBody(client, url, method.CONNECT)
}

//noinspection GoUnusedExportedFunction
func OPTIONS(client *http.Client, url string) (resp *http.Response, err error) {
	return noBody(client, url, method.OPTIONS)
}

//noinspection GoUnusedExportedFunction
func TRACE(client *http.Client, url string) (resp *http.Response, err error) {
	return noBody(client, url, method.TRACE)
}

func noBody(client *http.Client, url string, verb method.Method) (resp *http.Response, err error) {
	req, err := http.NewRequest(string(verb), url, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func bodied(client *http.Client, url, contentType string, body io.Reader, verb method.Method) (resp *http.Response, err error) {
	req, err := http.NewRequest(string(verb), url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return client.Do(req)
}
