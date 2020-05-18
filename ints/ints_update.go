package ints

//noinspection GoUnusedExportedFunction
func Update(curV, newV int, currentUpdated bool) (value int, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV int, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV int) bool {
	return curV != newV
}
