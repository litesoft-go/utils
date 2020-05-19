package int64s

//noinspection GoUnusedExportedFunction
func IsOneOf(check int64, acceptables ...int64) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
