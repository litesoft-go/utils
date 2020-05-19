package int16s

//noinspection GoUnusedExportedFunction
func Update(curV, newV int16, currentUpdated bool) (value int16, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV int16, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV int16) bool {
	return curV != newV
}
