package uint32s

//noinspection GoUnusedExportedFunction
func IsOneOf(check uint32, acceptables ...uint32) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
