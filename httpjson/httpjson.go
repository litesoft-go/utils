package httpjson

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Get attempts to perform an http GET method with the provided url.
// It returns the request response which includes the URL used, http Verb & StatusCode (-1 if unavailable),
// Response Data (assumed to be json, but could be empty, e.g. response not available), and
// an optional error.
//noinspection GoUnusedExportedFunction
func Get(client *http.Client, url string) *RequestResponse {
	return handle(client, url, "get", responseProviderGet)
}

// Post attempts to perform an http POST method with the provided url and post body (assumes "application/json").
// It returns the request response which includes the URL used, http Verb & StatusCode (-1 if unavailable),
// Response Data (assumed to be json, but could be empty, e.g. response not available), and
// an optional error.
//noinspection GoUnusedExportedFunction
func Post(client *http.Client, url, postBody string) *RequestResponse {
	return handle(client, url, "post",
		(&postResponseProvider{postBody: postBody}).post)
}

// Delete attempts to perform an http DELETE method with the provided url.
// It returns the request response which includes the URL used, http Verb & StatusCode (-1 if unavailable),
// Response Data (assumed to be json, but could be empty, e.g. response not available), and
// an optional error.
//noinspection GoUnusedExportedFunction
func Delete(client *http.Client, url string) *RequestResponse {
	return handle(client, url, "delete", responseProviderDelete)
}

type responseProvider func(client *http.Client, url string) (*http.Response, error)

func responseProviderDelete(client *http.Client, url string) (*http.Response, error) {
	// Create request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func responseProviderGet(client *http.Client, url string) (*http.Response, error) {
	return client.Get(url)
}

type postResponseProvider struct {
	postBody string
}

func (r *postResponseProvider) post(client *http.Client, url string) (*http.Response, error) {
	return client.Post(url,
		"application/json", strings.NewReader(r.postBody))
}

func handle(client *http.Client, url, httpMethod string, provider responseProvider) (result *RequestResponse) {
	result = &RequestResponse{Verb: httpMethod, URL: url, StatusCode: -1} // default!
	var err error
	defer func() {
		result.Err = err
	}()

	if client == nil {
		client = http.DefaultClient
	}
	var response *http.Response // to not use :=
	response, err = provider(client, url)
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
