package significant

import (
	"strings"
	// standard libs only above!
)

func OrEmpty(in string) string {
	if in == "" {
		return in
	}
	return strings.TrimSpace(in)
}

//noinspection GoUnusedExportedFunction
func Or(in, p2nd string, pOrs ...string) (rv string) {
	rv = OrEmpty(in)
	if rv == "" {
		rv = OrEmpty(p2nd)
		if rv == "" {
			for i := range pOrs {
				rv = OrEmpty(pOrs[i])
				if rv != "" {
					return
				}
			}
		}
	}
	return
}
