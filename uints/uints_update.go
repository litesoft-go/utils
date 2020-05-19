package uints

//noinspection GoUnusedExportedFunction
func Update(curV, newV uint, currentUpdated bool) (value uint, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV uint, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV uint) bool {
	return curV != newV
}
