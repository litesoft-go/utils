package types

const (
	UTF8charSet = "charset=UTF-8"

	JSONcontentType            = "application/json"
	JSONcontentTypeWithCharSet = JSONcontentType + "; " + UTF8charSet

	TEXTcontentType            = "text/html"
	TEXTcontentTypeWithCharSet = TEXTcontentType + "; " + UTF8charSet
)
