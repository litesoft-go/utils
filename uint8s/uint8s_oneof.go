package uint8s

//noinspection GoUnusedExportedFunction
func IsOneOf(check uint8, acceptables ...uint8) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
