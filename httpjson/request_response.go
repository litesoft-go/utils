package httpjson

import (
	"errors"
	"fmt"

	// standard libs only above!

	"github.com/litesoft-go/utils/strs"
)

type RequestResponse struct {
	Verb         string
	URL          string
	StatusCode   int
	ResponseData []byte
	Err          error
}

func (r *RequestResponse) BadRequest(verb, url, err string) *RequestResponse {
	rr := r
	if rr == nil {
		rr = &RequestResponse{}
	}
	rr.Verb = verb
	rr.URL = url
	rr.Err = errors.New(err)
	return rr
}

func (r *RequestResponse) OK() bool {
	return (r != nil) && // Left to Right!
		(r.StatusCode == 200) && (r.Err == nil)
}

func (r *RequestResponse) LikelyParsable() bool {
	return r.OK() && // Left to Right!
		(len(r.ResponseData) != 0)
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
