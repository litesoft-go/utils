package int32s

//noinspection GoUnusedExportedFunction
func Update(curV, newV int32, currentUpdated bool) (value int32, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV int32, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV int32) bool {
	return curV != newV
}
