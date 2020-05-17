package significant

import "strings"

func OrEmpty(in string) string {
	if in == "" {
		return in
	}
	return strings.TrimSpace(in)
}

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
