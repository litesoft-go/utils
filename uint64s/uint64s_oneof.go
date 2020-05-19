package uint64s

//noinspection GoUnusedExportedFunction
func IsOneOf(check uint64, acceptables ...uint64) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
