package strs

//noinspection GoUnusedExportedFunction
func Update(curV, newV string, currentUpdated bool) (value string, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV string, currentUpdated bool, setter Sync) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

func different(curV, newV string) bool {
	return curV != newV
}
