package uint32s

//noinspection GoUnusedExportedFunction
func Update(curV, newV uint32, currentUpdated bool) (value uint32, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV uint32, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV uint32) bool {
	return curV != newV
}
