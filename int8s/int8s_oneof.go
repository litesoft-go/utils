package int8s

//noinspection GoUnusedExportedFunction
func IsOneOf(check int8, acceptables ...int8) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
