package uint8s

//noinspection GoUnusedExportedFunction
func Update(curV, newV uint8, currentUpdated bool) (value uint8, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV uint8, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV uint8) bool {
	return curV != newV
}
