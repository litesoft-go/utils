package uint16s

//noinspection GoUnusedExportedFunction
func IsOneOf(check uint16, acceptables ...uint16) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
