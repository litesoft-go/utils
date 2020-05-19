package int16s

//noinspection GoUnusedExportedFunction
func IsOneOf(check int16, acceptables ...int16) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
