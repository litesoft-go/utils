package uint16s

//noinspection GoUnusedExportedFunction
func Update(curV, newV uint16, currentUpdated bool) (value uint16, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV uint16, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV uint16) bool {
	return curV != newV
}
