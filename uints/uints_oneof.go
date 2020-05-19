package uints

//noinspection GoUnusedExportedFunction
func IsOneOf(check uint, acceptables ...uint) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}
