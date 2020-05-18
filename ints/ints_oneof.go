package ints

//noinspection GoUnusedExportedFunction
func IsOneOf(check int, acceptables ...int) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
