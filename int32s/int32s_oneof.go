package int32s

//noinspection GoUnusedExportedFunction
func IsOneOf(check int32, acceptables ...int32) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
