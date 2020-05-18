package httpjson

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// standard libs only above!

	"github.com/litesoft-go/utils/strs"
)

var Client *http.Client

type RequestResponse struct {
	Verb         string
	URL          string
	StatusCode   int
	ResponseData []byte
	Err          error
}

func (r *RequestResponse) LikelyParsable() bool {
	return (r != nil) && // Left to Right!
		(r.StatusCode == 200) && (len(r.ResponseData) != 0) && (r.Err == nil)
}

func (r *RequestResponse) String() string {
	ib := strs.NewIndentedBuilder(4)
	r.IndentString("httpjson.RequestResponse", ib)
	return ib.String()
}

func (r *RequestResponse) IndentString(prefix string, ib *strs.IndentedBuilder) {
	if r == nil {
		ib.AddLine("No " + prefix)
	} else {
		ib.AddLine(fmt.Sprintf("%s - %s (%d): %s", prefix, r.Verb, r.StatusCode, r.URL))
		if r.Err != nil {
			ib.AddIndent()
			ib.AddLine("err=", r.Err.Error())
			ib.DropIndent()
		}
		if len(r.ResponseData) != 0 {
			ib.AddIndent()
			ib.AddLine("responseData:")
			ib.AddIndent()
			ib.AddLine(string(r.ResponseData))
			ib.DropIndent()
			ib.DropIndent()
		}
	}
}

// Get attempts to perform an http GET method with the provided url.
// It returns the responses status code (or -1 if unavailable),
// the response data (assumed to be json, but could be empty, e.g. response not available), and
// an optional error.
//noinspection GoUnusedExportedFunction
func Get(url string) *RequestResponse {
	provider := &getResponseProvider{}
	return handle(url, provider.get, "get")
}

// Get attempts to perform an http POST method with the provided url and post body (assumes "application/json").
// It returns the responses status code (or -1 if unavailable),
// the response data (assumed to be json, but could be empty, e.g. response not available), and
// an optional error.
//noinspection GoUnusedExportedFunction
func Post(url, postBody string) *RequestResponse {
	provider := &postResponseProvider{postBody: postBody}
	return handle(url, provider.post, "post")
}

type responseProvider func(url string) (*http.Response, error)

type getResponseProvider struct {
}

func (r *getResponseProvider) get(url string) (*http.Response, error) {
	return getClient().Get(url)
}

type postResponseProvider struct {
	postBody string
}

func (r *postResponseProvider) post(url string) (*http.Response, error) {
	return getClient().Post(url, "application/json", strings.NewReader(r.postBody))
}

func getClient() *http.Client {
	client := Client
	if client != nil {
		return client
	}
	return http.DefaultClient
}

func handle(url string, provider responseProvider, httpMethod string) (result *RequestResponse) {
	result = &RequestResponse{Verb: httpMethod, URL: url, StatusCode: -1} // default!
	var err error
	defer func() {
		result.Err = err
	}()
	var response *http.Response // to not use :=
	response, err = provider(url)
	if response != nil {
		defer func() {
			_ = response.Body.Close() // Per Dave Cheney 2017 - auto drains!
		}()
	}
	if err == nil {
		if response == nil {
			err = errors.New("no response")
		} else {
			result.StatusCode = response.StatusCode

			result.ResponseData, err = ioutil.ReadAll(response.Body)
		}
	}
	return
}
