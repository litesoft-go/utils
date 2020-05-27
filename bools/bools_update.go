package bools

//noinspection GoUnusedExportedFunction
func Update(curV, newV, currentUpdated bool) (value, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV bool) bool {
	return curV != newV
}
