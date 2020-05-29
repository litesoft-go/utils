package method

import "net/http"

type Method string

const (
	POST    Method = http.MethodPost
	PUT     Method = http.MethodPut
	PATCH   Method = http.MethodPatch
	GET     Method = http.MethodGet
	HEAD    Method = http.MethodHead
	DELETE  Method = http.MethodDelete
	CONNECT Method = http.MethodConnect
	OPTIONS Method = http.MethodOptions
	TRACE   Method = http.MethodTrace
)
