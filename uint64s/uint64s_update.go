package uint64s

//noinspection GoUnusedExportedFunction
func Update(curV, newV uint64, currentUpdated bool) (value uint64, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV uint64, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV uint64) bool {
	return curV != newV
}
