package int64s

//noinspection GoUnusedExportedFunction
func Update(curV, newV int64, currentUpdated bool) (value int64, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV int64, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV int64) bool {
	return curV != newV
}
